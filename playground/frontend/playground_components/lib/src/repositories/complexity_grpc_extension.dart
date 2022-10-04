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

import '../api/v1/api.pbgrpc.dart' as grpc;
import '../enums/complexity.dart';

extension GrpcComplecity on grpc.Complexity {
  Complexity get model {
    switch (this) {
      case grpc.Complexity.COMPLEXITY_BASIC:
        return Complexity.basic;
      case grpc.Complexity.COMPLEXITY_MEDIUM:
        return Complexity.medium;
      case grpc.Complexity.COMPLEXITY_ADVANCED:
        return Complexity.advanced;
      case grpc.Complexity.COMPLEXITY_UNSPECIFIED:
        return Complexity.unspecified;
    }
    throw Exception('Unknown complexity: $this');
  }
}
