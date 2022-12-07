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

package org.apache.beam.examples;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.coders.*;
import org.apache.beam.sdk.io.kafka.KafkaIO;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.schemas.JavaFieldSchema;
import org.apache.beam.sdk.schemas.Schema;
import org.apache.beam.sdk.schemas.annotations.DefaultSchema;
import org.apache.beam.sdk.schemas.annotations.SchemaCreate;
import org.apache.beam.sdk.schemas.transforms.Group;
import org.apache.beam.sdk.transforms.*;
import org.apache.beam.sdk.util.StreamUtils;
import org.apache.beam.sdk.values.Row;
import org.apache.beam.sdk.values.TypeDescriptor;
import org.apache.kafka.common.TopicPartition;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.kafka.common.serialization.StringDeserializer;

// beam-playground:
//   name: ZookeeperlessKafkaNycTaxiPassengerCountJson
//   description: Test example with Apache Kafka
//   multifile: false
//   context_line: 55
//   categories:
//     - Filtering
//     - Options
//     - Quickstart
//   complexity: MEDIUM
//   tags:
//     - filter
//     - strings
//     - emulator
//   emulators:
//      kafka:
//          topic:
//              id: dataset
//              dataset: NYCTaxi1000
//   datasets:
//      NYCTaxi1000:
//          location: local
//          format: json


public class ZookeeperlessKafkaNycTaxiPassengerCountJson {
    public static void main(String[] args) {
        System.out.println("Hello");
    }}
/*        PipelineOptions options = PipelineOptionsFactory.fromArgs(args).create();
        Pipeline p = Pipeline.create(options);
        CoderProvider stringCustomProvider = CoderProviders.fromStaticMethods(VendorToPassengerDTOArray.class, CustomArrayCoder.class);
        p.getCoderRegistry().registerCoderProvider(stringCustomProvider);
        final Schema schema = Schema.builder()
                .addInt32Field("passengerCount")
                .addInt32Field("vendorID")
                .build();

        Map<String, Object> consumerConfig = new HashMap<>();
        consumerConfig.put("auto.offset.reset", "earliest");
        ObjectMapper om = new ObjectMapper();

        p
                .apply("ReadFromKafka", KafkaIO.<String, String>read()
                        .withBootstrapServers("kafka_server:9092") // The argument is predefined to a correct value. Do not change it manually.
                        .withTopicPartitions(Collections.singletonList(new TopicPartition("dataset", 0))) // The argument is predefined to a correct value. Do not change it manually.
                        .withKeyDeserializer(StringDeserializer.class)
                        .withValueDeserializer(StringDeserializer.class)
                        .withConsumerConfigUpdates(consumerConfig)
                        .withMaxNumRecords(1000)
                        .withoutMetadata())
                .apply("CreateValues", Values.create())
                .apply("ExtractData", ParDo.of(new DoFn<String, VendorToPassengerDTOArray>() {
                    @ProcessElement
                    @SuppressWarnings("unused")
                    public void processElement(ProcessContext c) throws JsonProcessingException {
                        final List<VendorToPassengerDTO> result = om.readValue(c.element(), new TypeReference<List<VendorToPassengerDTO>>() {
                        });
                        final VendorToPassengerDTOArray dtOs = new VendorToPassengerDTOArray(result);
                        c.output(dtOs);
                    }
                }))
                .apply("FlatMap", FlatMapElements.into(TypeDescriptor.of(VendorToPassengerDTO.class)).via(x -> x.data))
                .setCoder(CustomCoder.of())
                .apply(MapElements.into(TypeDescriptor.of(Object.class)).via(it -> it))
                .setSchema(schema, TypeDescriptor.of(Object.class), input -> {
                    VendorToPassengerDTO dto = (VendorToPassengerDTO) input;
                    return Row.withSchema(schema)
                            .addValues(dto.passengerCount, dto.vendorID)
                            .build();
                }, input -> new VendorToPassengerDTO(input.getInt32("passengerCount"), input.getInt32("vendorID")))
                .apply(Group.byFieldNames("vendorID").aggregateField("passengerCount", Sum.ofIntegers(), "totalPassengerCount"))
                .apply("CheckResult", ParDo.of(new LogOutput<>()));

        p.run().waitUntilFinish();
    }
}

class CustomCoder extends Coder<VendorToPassengerDTO> {

    final ObjectMapper om = new ObjectMapper();

    private static final CustomCoder INSTANCE = new CustomCoder();

    public static CustomCoder of() {
        return INSTANCE;
    }

    @Override
    public void encode(VendorToPassengerDTO dto, OutputStream outStream) throws IOException {
        final String result = dto.toString();
        outStream.write(result.getBytes());
    }

    @Override
    public VendorToPassengerDTO decode(InputStream inStream) throws IOException {
        final String serializedDTOs = new String(StreamUtils.getBytesWithoutClosing(inStream));
        return om.readValue(serializedDTOs, VendorToPassengerDTO.class);
    }

    @Override
    public List<? extends Coder<?>> getCoderArguments() {
        return Collections.emptyList();
    }

    @Override
    public void verifyDeterministic() {

    }
}

class CustomArrayCoder extends Coder<VendorToPassengerDTOArray> {

    private static final CustomArrayCoder INSTANCE = new CustomArrayCoder();

    public static CustomArrayCoder of() {
        return INSTANCE;
    }

    final ObjectMapper om = new ObjectMapper();

    @Override
    public void encode(VendorToPassengerDTOArray data, OutputStream outStream) throws IOException {
        final String result = data.toString();
        outStream.write(result.getBytes());
    }

    @Override
    public VendorToPassengerDTOArray decode(final InputStream inStream) throws IOException {
        final String serializedDTOs = new String(StreamUtils.getBytesWithoutClosing(inStream));
        final List<VendorToPassengerDTO> result = om.readValue(serializedDTOs, new TypeReference<List<VendorToPassengerDTO>>() {
        });
        return new VendorToPassengerDTOArray(result);
    }

    @Override
    public List<? extends Coder<?>> getCoderArguments() {
        return Collections.emptyList();
    }

    @Override
    public void verifyDeterministic() {
    }
}

class VendorToPassengerDTOArray {
    List<VendorToPassengerDTO> data;

    public VendorToPassengerDTOArray(List<VendorToPassengerDTO> data) {
        this.data = data;
    }

    @Override
    public String toString() {
        final StringBuilder strBuilder = new StringBuilder("[");
        int lastElement = data.size() - 1;
        for (int i = 0; i < data.size(); i++) {
            if (i == lastElement) {
                strBuilder.append(String.format("{\"PassengerCount\":%d,\"VendorID\":%d}]", data.get(i).getPassengerCount(), data.get(i).getVendorID()));
                break;
            }
            strBuilder.append(String.format("{\"PassengerCount\":%d,\"VendorID\":%d},", data.get(i).getPassengerCount(), data.get(i).getVendorID()));
        }
        return strBuilder.toString();
    }
}

@DefaultSchema(JavaFieldSchema.class)
class VendorToPassengerDTO {
    @JsonProperty(value = "PassengerCount")
    Integer passengerCount;
    @JsonProperty(value = "VendorID")
    Integer vendorID;

    @SchemaCreate
    public VendorToPassengerDTO(Integer passengerCount, Integer vendorID) {
        this.passengerCount = passengerCount;
        this.vendorID = vendorID;
    }

    public Integer getVendorID() {
        return this.vendorID;
    }

    public Integer getPassengerCount() {
        return this.passengerCount;
    }

    @Override
    public String toString() {
        return String.format("{\"PassengerCount\":%d,\"VendorID\":%d}", passengerCount, vendorID);
    }
}

class LogOutput<T> extends DoFn<T, T> {

    private final String prefix;

    LogOutput() {
        this.prefix = "Processing element";
    }

    @ProcessElement
    public void processElement(ProcessContext c) {
        System.out.println(prefix + ": " + c.element());
    }
}*/
