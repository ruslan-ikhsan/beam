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

resource "google_cloudbuild_trigger" "backend_builder" {
  name = "backend-builder"
  description = "Builds the base image used for Beam Playground backend builds"
  github {
    owner = var.github_repository_owner
    name  = var.github_repository_name
    push {
      branch = var.github_repository_branch
    }
  }
  // Disabled because we only want to run it manually
  disabled = true
  build {
    step {
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t",
        "${local.backend_playground_repository_uri_prefix}-builder:${var.image_tag}",
        "-f",
        "playground/infrastructure/02.build/Dockerfile-backend-builder",
        "."
      ]
    }
    images = [
      "${local.backend_playground_repository_uri_prefix}-builder:${var.image_tag}"
    ]
  }
}

resource "google_cloudbuild_trigger" "backend" {
  for_each = local.containers.backend
  name = "build-backend-${each.key}"
  description = "Builds the Beam Playground ${each.key} docker container. (Requires: ${local.backend_playground_repository_uri_prefix}-builder image)"
  github {
    owner = var.github_repository_owner
    name  = var.github_repository_name
    push {
      branch = var.github_repository_branch
    }
  }
  included_files = [
    "playground/backend/**"
  ]

  build {
    step {
      name = "gcr.io/cloud-builders/docker"
      entrypoint = "./gradlew"
      args = [
        ":playground:backend:containers:${each.key}:docker",
        "--stacktrace"
      ]
    }
    images = [
      "${local.backend_playground_repository_uri_prefix}-${each.key}:${var.image_tag}"
    ]
  }
}