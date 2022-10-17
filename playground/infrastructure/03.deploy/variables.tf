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

variable "application_name" {
  type = string
  description = "The application name used as a common resource name"
  default = "beam-playground"
}

variable "artifact_registry_repository_id" {
  type = string
  description = "The ID of the Google Cloud Artifact Registry"
  default = "beam-playground"
}

variable "image_tag" {
  type = string
  description = "The tag to apply to image builds"
  default = "latest"
}

variable "memory_store_configuration" {
  // See https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/redis_instance#argument-reference
  type = object({
    memory_size_gb = number
    tier = string
    replica_count = number
    read_replicas_mode = string
    redis_version = string
  })
  default = {
    memory_size_gb = 5
    tier = "STANDARD_HA"
    replica_count = 1
    read_replicas_mode = "READ_REPLICAS_ENABLED"
    redis_version = "REDIS_6_X"
  }
}

variable "project" {
  type = string
  description = "The ID of the Google Cloud project within which resources are provisioned"
}

variable "region" {
  type = string
  description = "The Google Cloud Platform (GCP) region in which to provision resources"
  default = "us-central1"
}

variable "service_account_id" {
  type = string
  description = "The application service account ID bound to application workloads"
  default = "beam-playground"
}

variable "subnetwork_cidr" {
  type = string
  description = "The address range for this subnet, in CIDR notation. Use a standard private VPC network address range: for example, 10.0.0.0/9."
  default = "10.128.0.0/20"
}

variable "vpc_access_connector_ip_range" {
  type        = string
  description = "The address range for the VPC access connector ip, in CIDR notation"
  default     = "10.8.0.0/28"
}