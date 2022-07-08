/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.apache.beam.fn.harness.data;

import com.google.auto.value.AutoValue;
import java.io.Closeable;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;
import org.apache.beam.fn.harness.HandlesSplits;
import org.apache.beam.fn.harness.control.BundleProgressReporter;
import org.apache.beam.fn.harness.control.Metrics;
import org.apache.beam.fn.harness.control.Metrics.BundleCounter;
import org.apache.beam.fn.harness.control.Metrics.BundleDistribution;
import org.apache.beam.model.fnexecution.v1.BeamFnApi.ProcessBundleDescriptor;
import org.apache.beam.model.pipeline.v1.RunnerApi;
import org.apache.beam.runners.core.construction.RehydratedComponents;
import org.apache.beam.runners.core.metrics.ExecutionStateTracker;
import org.apache.beam.runners.core.metrics.MetricsContainerStepMap;
import org.apache.beam.runners.core.metrics.MonitoringInfoConstants;
import org.apache.beam.runners.core.metrics.MonitoringInfoConstants.Labels;
import org.apache.beam.runners.core.metrics.MonitoringInfoConstants.TypeUrns;
import org.apache.beam.runners.core.metrics.MonitoringInfoConstants.Urns;
import org.apache.beam.runners.core.metrics.MonitoringInfoMetricName;
import org.apache.beam.runners.core.metrics.ShortIdMap;
import org.apache.beam.runners.core.metrics.SimpleExecutionState;
import org.apache.beam.runners.core.metrics.SimpleMonitoringInfoBuilder;
import org.apache.beam.runners.core.metrics.SimpleStateRegistry;
import org.apache.beam.sdk.coders.Coder;
import org.apache.beam.sdk.fn.data.FnDataReceiver;
import org.apache.beam.sdk.metrics.Distribution;
import org.apache.beam.sdk.metrics.MetricsContainer;
import org.apache.beam.sdk.metrics.MetricsEnvironment;
import org.apache.beam.sdk.util.WindowedValue;
import org.apache.beam.sdk.util.WindowedValue.WindowedValueCoder;
import org.apache.beam.sdk.util.common.ElementByteSizeObserver;
import org.apache.beam.vendor.grpc.v1p43p2.com.google.protobuf.ByteString;
import org.apache.beam.vendor.guava.v26_0_jre.com.google.common.collect.ImmutableList;

/**
 * The {@code PCollectionConsumerRegistry} is used to maintain a collection of consuming
 * FnDataReceiver for each pCollectionId. Registering with this class allows inserting an element
 * count counter for every pCollection. A combined MultiplexingConsumer (Wrapped with an
 * ElementCountFnDataReceiver) is returned by calling getMultiplexingConsumer.
 */
@SuppressWarnings({
  "rawtypes", // TODO(https://github.com/apache/beam/issues/20447)
  "nullness" // TODO(https://github.com/apache/beam/issues/20497)
})
public class PCollectionConsumerRegistry {

  /** Stores metadata about each consumer so that the appropriate metrics tracking can occur. */
  @AutoValue
  @AutoValue.CopyAnnotations
  @SuppressWarnings({"rawtypes"})
  abstract static class ConsumerAndMetadata {
    public static ConsumerAndMetadata forConsumer(
        FnDataReceiver consumer,
        String pTransformId,
        SimpleExecutionState state,
        MetricsContainer metricsContainer) {
      return new AutoValue_PCollectionConsumerRegistry_ConsumerAndMetadata(
          consumer, pTransformId, state, metricsContainer);
    }

    public abstract FnDataReceiver getConsumer();

    public abstract String getPTransformId();

    public abstract SimpleExecutionState getExecutionState();

    public abstract MetricsContainer getMetricsContainer();
  }

  private final MetricsContainerStepMap metricsContainerRegistry;
  private final ExecutionStateTracker stateTracker;
  private final ShortIdMap shortIdMap;
  private final Map<String, List<ConsumerAndMetadata>> pCollectionIdsToConsumers;
  private final Map<String, FnDataReceiver> pCollectionIdsToWrappedConsumer;
  private final SimpleStateRegistry executionStates;
  private final BundleProgressReporter.Registrar bundleProgressReporterRegistrar;
  private final ProcessBundleDescriptor processBundleDescriptor;
  private final RehydratedComponents rehydratedComponents;

  public PCollectionConsumerRegistry(
      MetricsContainerStepMap metricsContainerRegistry,
      ExecutionStateTracker stateTracker,
      ShortIdMap shortIdMap,
      BundleProgressReporter.Registrar bundleProgressReporterRegistrar,
      ProcessBundleDescriptor processBundleDescriptor) {
    this.metricsContainerRegistry = metricsContainerRegistry;
    this.stateTracker = stateTracker;
    this.shortIdMap = shortIdMap;
    this.pCollectionIdsToConsumers = new HashMap<>();
    this.pCollectionIdsToWrappedConsumer = new HashMap<>();
    this.executionStates = new SimpleStateRegistry();
    this.bundleProgressReporterRegistrar = bundleProgressReporterRegistrar;
    this.processBundleDescriptor = processBundleDescriptor;
    this.rehydratedComponents =
        RehydratedComponents.forComponents(
            RunnerApi.Components.newBuilder()
                .putAllCoders(processBundleDescriptor.getCodersMap())
                .putAllPcollections(processBundleDescriptor.getPcollectionsMap())
                .putAllWindowingStrategies(processBundleDescriptor.getWindowingStrategiesMap())
                .build());
  }

  /**
   * Register the specified consumer to handle the elements in the pCollection associated with
   * pCollectionId. All consumers must be registered before extracting the combined consumer by
   * calling getMultiplexingConsumer(), or an exception will be thrown.
   *
   * <p>This will cause both Element Count and Process Bundle Execution time metrics to be
   * collected.
   *
   * @param pCollectionId
   * @param pTransformId
   * @param consumer
   * @param <T> the element type of the PCollection
   * @throws RuntimeException if {@code register()} is called after {@code
   *     getMultiplexingConsumer()} is called.
   */
  public <T> void register(
      String pCollectionId, String pTransformId, FnDataReceiver<WindowedValue<T>> consumer) {
    // Just save these consumers for now, but package them up later with an
    // ElementCountFnDataReceiver and possibly a MultiplexingFnDataReceiver
    // if there are multiple consumers.
    if (pCollectionIdsToWrappedConsumer.containsKey(pCollectionId)) {
      throw new RuntimeException(
          "New consumers for a pCollectionId cannot be register()-d after "
              + "calling getMultiplexingConsumer.");
    }

    HashMap<String, String> labelsMetadata = new HashMap<>();
    labelsMetadata.put(MonitoringInfoConstants.Labels.PTRANSFORM, pTransformId);
    SimpleExecutionState state =
        new SimpleExecutionState(
            ExecutionStateTracker.PROCESS_STATE_NAME,
            MonitoringInfoConstants.Urns.PROCESS_BUNDLE_MSECS,
            labelsMetadata);
    executionStates.register(state);

    List<ConsumerAndMetadata> consumerAndMetadatas =
        pCollectionIdsToConsumers.computeIfAbsent(pCollectionId, (unused) -> new ArrayList<>());
    consumerAndMetadatas.add(
        ConsumerAndMetadata.forConsumer(
            consumer, pTransformId, state, metricsContainerRegistry.getContainer(pTransformId)));
  }

  /**
   * New consumers should not be register()-ed after calling this method. This will cause a
   * RuntimeException, as this would fail to properly wrap the late-added consumer to track metrics.
   *
   * @return A {@link FnDataReceiver} which directly wraps all the registered consumers.
   */
  public FnDataReceiver<WindowedValue<?>> getMultiplexingConsumer(String pCollectionId) {
    return pCollectionIdsToWrappedConsumer.computeIfAbsent(
        pCollectionId,
        pcId -> {
          if (!processBundleDescriptor.containsPcollections(pCollectionId)) {
            throw new IllegalArgumentException(
                String.format("Unknown PCollection id %s", pCollectionId));
          }
          String coderId =
              processBundleDescriptor.getPcollectionsOrThrow(pCollectionId).getCoderId();
          Coder<?> coder;
          try {
            Coder<?> maybeWindowedValueInputCoder = rehydratedComponents.getCoder(coderId);
            // TODO: Stop passing windowed value coders within PCollections.
            if (maybeWindowedValueInputCoder instanceof WindowedValue.WindowedValueCoder) {
              coder = ((WindowedValueCoder) maybeWindowedValueInputCoder).getValueCoder();
            } else {
              coder = maybeWindowedValueInputCoder;
            }
          } catch (IOException e) {
            throw new IllegalStateException(
                String.format("Unable to materialize coder %s", coderId), e);
          }
          List<ConsumerAndMetadata> consumerAndMetadatas =
              pCollectionIdsToConsumers.computeIfAbsent(
                  pCollectionId, (unused) -> new ArrayList<>());

          if (consumerAndMetadatas.size() == 1) {
            ConsumerAndMetadata consumerAndMetadata = consumerAndMetadatas.get(0);
            if (consumerAndMetadata.getConsumer() instanceof HandlesSplits) {
              return new SplittingMetricTrackingFnDataReceiver(pcId, coder, consumerAndMetadata);
            }
            return new MetricTrackingFnDataReceiver(pcId, coder, consumerAndMetadata);
          } else {
            /* TODO(SDF), Consider supporting splitting each consumer individually. This would never
            come up in the existing SDF expansion, but might be useful to support fused SDF nodes.
            This would require dedicated delivery of the split results to each of the consumers
            separately. */
            return new MultiplexingMetricTrackingFnDataReceiver(
                pcId, coder, ImmutableList.copyOf(consumerAndMetadatas));
          }
        });
  }

  /** @return Execution Time Monitoring data based on the tracked start or finish function. */
  public Map<String, ByteString> getExecutionTimeMonitoringData(ShortIdMap shortIds) {
    return executionStates.getExecutionTimeMonitoringData(shortIds);
  }

  /** Reset the execution states of the registered functions. */
  public void reset() {
    executionStates.reset();
  }

  /**
   * A wrapping {@code FnDataReceiver<WindowedValue<T>>} which counts the number of elements
   * consumed by the original {@code FnDataReceiver<WindowedValue<T>> consumer} and sets up metrics
   * for tracking PTransform processing time.
   *
   * @param <T> - The receiving type of the PTransform.
   */
  private class MetricTrackingFnDataReceiver<T> implements FnDataReceiver<WindowedValue<T>> {
    private final FnDataReceiver<WindowedValue<T>> delegate;
    private final SimpleExecutionState state;
    private final BundleCounter elementCountCounter;
    private final SampleByteSizeDistribution<T> sampledByteSizeDistribution;
    private final Coder<T> coder;
    private final MetricsContainer metricsContainer;

    public MetricTrackingFnDataReceiver(
        String pCollectionId, Coder<T> coder, ConsumerAndMetadata consumerAndMetadata) {
      this.delegate = consumerAndMetadata.getConsumer();
      this.state = consumerAndMetadata.getExecutionState();

      HashMap<String, String> labels = new HashMap<>();
      labels.put(Labels.PCOLLECTION, pCollectionId);
      MonitoringInfoMetricName elementCountMetricName =
          MonitoringInfoMetricName.named(MonitoringInfoConstants.Urns.ELEMENT_COUNT, labels);
      String elementCountShortId =
          shortIdMap.getOrCreateShortId(
              new SimpleMonitoringInfoBuilder()
                  .setUrn(MonitoringInfoConstants.Urns.ELEMENT_COUNT)
                  .setType(TypeUrns.SUM_INT64_TYPE)
                  .setLabels(labels)
                  .build());
      this.elementCountCounter =
          Metrics.bundleProcessingThreadCounter(elementCountShortId, elementCountMetricName);
      bundleProgressReporterRegistrar.register(elementCountCounter);

      MonitoringInfoMetricName sampledByteSizeMetricName =
          MonitoringInfoMetricName.named(Urns.SAMPLED_BYTE_SIZE, labels);
      String sampledByteSizeShortId =
          shortIdMap.getOrCreateShortId(
              new SimpleMonitoringInfoBuilder()
                  .setUrn(Urns.SAMPLED_BYTE_SIZE)
                  .setType(TypeUrns.DISTRIBUTION_INT64_TYPE)
                  .setLabels(labels)
                  .build());
      BundleDistribution sampledByteSizeUnderlyingDistribution =
          Metrics.bundleProcessingThreadDistribution(
              sampledByteSizeShortId, sampledByteSizeMetricName);
      this.sampledByteSizeDistribution =
          new SampleByteSizeDistribution<>(sampledByteSizeUnderlyingDistribution);
      bundleProgressReporterRegistrar.register(sampledByteSizeUnderlyingDistribution);

      this.coder = coder;
      this.metricsContainer = consumerAndMetadata.getMetricsContainer();
    }

    @Override
    public void accept(WindowedValue<T> input) throws Exception {
      // Increment the counter for each window the element occurs in.
      this.elementCountCounter.inc(input.getWindows().size());
      // TODO(https://github.com/apache/beam/issues/20730): Consider updating size per window when
      // we have window optimization.
      this.sampledByteSizeDistribution.tryUpdate(input.getValue(), this.coder);

      // Wrap the consumer with extra logic to set the metric container with the appropriate
      // PTransform context. This ensures that user metrics obtain the pTransform ID when they are
      // created. Also use the ExecutionStateTracker and enter an appropriate state to track the
      // Process Bundle Execution time metric.
      try (Closeable closeable = MetricsEnvironment.scopedMetricsContainer(metricsContainer)) {
        try (Closeable trackerCloseable = stateTracker.enterState(state)) {
          this.delegate.accept(input);
        }
      }
      this.sampledByteSizeDistribution.finishLazyUpdate();
    }
  }

  /**
   * A wrapping {@code FnDataReceiver<WindowedValue<T>>} which counts the number of elements
   * consumed by the original {@code FnDataReceiver<WindowedValue<T>> consumers} and sets up metrics
   * for tracking PTransform processing time.
   *
   * @param <T> - The receiving type of the PTransform.
   */
  private class MultiplexingMetricTrackingFnDataReceiver<T>
      implements FnDataReceiver<WindowedValue<T>> {
    private final List<ConsumerAndMetadata> consumerAndMetadatas;
    private final BundleCounter elementCountCounter;
    private final SampleByteSizeDistribution<T> sampledByteSizeDistribution;
    private final Coder<T> coder;

    public MultiplexingMetricTrackingFnDataReceiver(
        String pCollectionId, Coder<T> coder, List<ConsumerAndMetadata> consumerAndMetadatas) {
      this.consumerAndMetadatas = consumerAndMetadatas;

      HashMap<String, String> labels = new HashMap<>();
      labels.put(Labels.PCOLLECTION, pCollectionId);
      MonitoringInfoMetricName elementCountMetricName =
          MonitoringInfoMetricName.named(MonitoringInfoConstants.Urns.ELEMENT_COUNT, labels);
      String elementCountShortId =
          shortIdMap.getOrCreateShortId(
              new SimpleMonitoringInfoBuilder()
                  .setUrn(MonitoringInfoConstants.Urns.ELEMENT_COUNT)
                  .setType(TypeUrns.SUM_INT64_TYPE)
                  .setLabels(labels)
                  .build());
      this.elementCountCounter =
          Metrics.bundleProcessingThreadCounter(elementCountShortId, elementCountMetricName);
      bundleProgressReporterRegistrar.register(elementCountCounter);

      MonitoringInfoMetricName sampledByteSizeMetricName =
          MonitoringInfoMetricName.named(Urns.SAMPLED_BYTE_SIZE, labels);
      String sampledByteSizeShortId =
          shortIdMap.getOrCreateShortId(
              new SimpleMonitoringInfoBuilder()
                  .setUrn(Urns.SAMPLED_BYTE_SIZE)
                  .setType(TypeUrns.DISTRIBUTION_INT64_TYPE)
                  .setLabels(labels)
                  .build());
      BundleDistribution sampledByteSizeUnderlyingDistribution =
          Metrics.bundleProcessingThreadDistribution(
              sampledByteSizeShortId, sampledByteSizeMetricName);
      this.sampledByteSizeDistribution =
          new SampleByteSizeDistribution<>(sampledByteSizeUnderlyingDistribution);
      bundleProgressReporterRegistrar.register(sampledByteSizeUnderlyingDistribution);

      this.coder = coder;
    }

    @Override
    public void accept(WindowedValue<T> input) throws Exception {
      // Increment the counter for each window the element occurs in.
      this.elementCountCounter.inc(input.getWindows().size());
      // TODO(https://github.com/apache/beam/issues/20730): Consider updating size per window
      // when we have window optimization.
      this.sampledByteSizeDistribution.tryUpdate(input.getValue(), coder);

      // Wrap the consumer with extra logic to set the metric container with the appropriate
      // PTransform context. This ensures that user metrics obtain the pTransform ID when they are
      // created. Also use the ExecutionStateTracker and enter an appropriate state to track the
      // Process Bundle Execution time metric.
      for (ConsumerAndMetadata consumerAndMetadata : consumerAndMetadatas) {

        try (Closeable closeable =
            MetricsEnvironment.scopedMetricsContainer(consumerAndMetadata.getMetricsContainer())) {
          try (Closeable trackerCloseable =
              stateTracker.enterState(consumerAndMetadata.getExecutionState())) {
            consumerAndMetadata.getConsumer().accept(input);
          }
        }
        this.sampledByteSizeDistribution.finishLazyUpdate();
      }
    }
  }

  /**
   * A wrapping {@code FnDataReceiver<WindowedValue<T>>} which counts the number of elements
   * consumed by the original {@code FnDataReceiver<WindowedValue<T>> consumer} and forwards split
   * and progress requests to the original consumer.
   *
   * @param <T> - The receiving type of the PTransform.
   */
  private class SplittingMetricTrackingFnDataReceiver<T> extends MetricTrackingFnDataReceiver<T>
      implements HandlesSplits {
    private final HandlesSplits delegate;

    public SplittingMetricTrackingFnDataReceiver(
        String pCollection, Coder<T> coder, ConsumerAndMetadata consumerAndMetadata) {
      super(pCollection, coder, consumerAndMetadata);
      this.delegate = (HandlesSplits) consumerAndMetadata.getConsumer();
    }

    @Override
    public SplitResult trySplit(double fractionOfRemainder) {
      return delegate.trySplit(fractionOfRemainder);
    }

    @Override
    public double getProgress() {
      return delegate.getProgress();
    }
  }

  private static class SampleByteSizeDistribution<T> {
    /** Basic implementation of {@link ElementByteSizeObserver} for use in size estimation. */
    private static class ByteSizeObserver extends ElementByteSizeObserver {
      private long observedSize = 0;

      @Override
      protected void reportElementSize(long elementSize) {
        observedSize += elementSize;
      }
    }

    final Distribution distribution;
    ByteSizeObserver byteCountObserver;

    public SampleByteSizeDistribution(Distribution distribution) {
      this.distribution = distribution;
      this.byteCountObserver = null;
    }

    public void tryUpdate(T value, Coder<T> coder) throws Exception {
      if (shouldSampleElement()) {
        // First try using byte size observer
        byteCountObserver = new ByteSizeObserver();
        coder.registerByteSizeObserver(value, byteCountObserver);

        if (!byteCountObserver.getIsLazy()) {
          byteCountObserver.advance();
          this.distribution.update(byteCountObserver.observedSize);
        }
      } else {
        byteCountObserver = null;
      }
    }

    public void finishLazyUpdate() {
      // Advance lazy ElementByteSizeObservers, if any.
      if (byteCountObserver != null && byteCountObserver.getIsLazy()) {
        byteCountObserver.advance();
        this.distribution.update(byteCountObserver.observedSize);
      }
    }

    private static final int RESERVOIR_SIZE = 10;
    private static final int SAMPLING_THRESHOLD = 30;
    private long samplingToken = 0;
    private long nextSamplingToken = 0;
    private Random randomGenerator = new Random();

    private boolean shouldSampleElement() {
      // Sampling probability decreases as the element count is increasing.
      // We unconditionally sample the first samplingCutoff elements. Calculating
      // nextInt(samplingToken) for each element is expensive, so after a threshold, calculate the
      // gap to next sample.
      // https://erikerlandson.github.io/blog/2015/11/20/very-fast-reservoir-sampling/

      // Reset samplingToken if it's going to exceed the max value.
      if (samplingToken + 1 == Long.MAX_VALUE) {
        samplingToken = 0;
        nextSamplingToken = getNextSamplingToken(samplingToken);
      }

      samplingToken++;
      // Use traditional sampling until the threshold of 30
      if (nextSamplingToken == 0) {
        if (samplingToken <= RESERVOIR_SIZE
            || randomGenerator.nextInt((int) samplingToken) < RESERVOIR_SIZE) {
          if (samplingToken > SAMPLING_THRESHOLD) {
            nextSamplingToken = getNextSamplingToken(samplingToken);
          }
          return true;
        }
      } else if (samplingToken >= nextSamplingToken) {
        nextSamplingToken = getNextSamplingToken(samplingToken);
        return true;
      }
      return false;
    }

    private long getNextSamplingToken(long samplingToken) {
      double gap =
          Math.log(1.0 - randomGenerator.nextDouble())
              / Math.log(1.0 - RESERVOIR_SIZE / (double) samplingToken);
      return samplingToken + (int) gap;
    }
  }
}
