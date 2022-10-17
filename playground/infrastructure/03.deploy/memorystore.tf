# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

resource "google_redis_instance" "cache" {
  depends_on = [
    google_service_networking_connection.private_service_connection,
  ]
  name           = var.application_name
  region = var.region
  display_name = var.application_name
  memory_size_gb = var.memory_store_configuration.memory_size_gb
  tier = var.memory_store_configuration.tier
  replica_count = var.memory_store_configuration.replica_count
  authorized_network = google_compute_network.default.name
  redis_version = var.memory_store_configuration.redis_version
  read_replicas_mode = var.memory_store_configuration.read_replicas_mode
  connect_mode = "PRIVATE_SERVICE_ACCESS"
}