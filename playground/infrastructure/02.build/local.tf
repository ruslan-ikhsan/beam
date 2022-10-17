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

locals {
  containers = {
    backend = toset([
      "go",
      "java",
      "python",
      "router",
      "scio",
    ])
  }
  artifact_registry_docker_prefix = "${var.region}-docker.pkg.dev"
  playground_repository_uri_prefix = "${local.artifact_registry_docker_prefix}/${var.project}/${var.artifact_registry_repository_id}"
  backend_playground_repository_uri_prefix = "${local.playground_repository_uri_prefix}/backend"
  backend_builder_uri = "${local.backend_playground_repository_uri_prefix}-builder:${var.image_tag}"
  frontend_builder_uri = "${local.playground_repository_uri_prefix}/frontend-builder:${var.image_tag}"
}