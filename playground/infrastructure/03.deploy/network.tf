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

resource "google_compute_network" "default" {
  name = var.application_name
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name = var.application_name
  network = google_compute_network.default.id
  region = var.region
  ip_cidr_range = var.subnetwork_cidr
}

resource "google_compute_global_address" "redis_service_range" {
  name          = "${var.application_name}-redis-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "private_service_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [
    google_compute_global_address.redis_service_range.name
  ]
}

resource "google_vpc_access_connector" "default" {
  name          = var.application_name
  network       = google_compute_network.default.id
  ip_cidr_range = var.vpc_access_connector_ip_range
}