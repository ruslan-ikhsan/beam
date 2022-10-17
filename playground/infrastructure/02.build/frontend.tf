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

resource "google_cloudbuild_trigger" "frontend_builder" {
  name = "frontend-builder"
  description = "Builds the base image used for Beam Playground frontend builds"
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
    timeout = "1800s"
    step {
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t",
        local.frontend_builder_uri,
        "-f",
        "playground/infrastructure/02.build/Dockerfile-frontend-builder",
        "."
      ]
    }
    images = [
      local.frontend_builder_uri,
    ]
  }
}

resource "google_cloudbuild_trigger" "frontend" {
  name = "build-frontend"
  description = "Builds the Beam Playground frontend docker container. (Requires: ${local.frontend_builder_uri} image)"
  github {
    owner = var.github_repository_owner
    name  = var.github_repository_name
    push {
      branch = var.github_repository_branch
    }
  }
  included_files = [
    "playground/frontend/**"
  ]

  build {
    timeout = "1800s"
    step {
      name = "gcr.io/cloud-builders/docker"
      dir = "playground/frontend"
      args = [
        "build",
        "-t",
        local.frontend_uri,
        "--build-arg=BUILDER=${local.frontend_builder_uri}",
        "."
      ]
    }
    images = [
      local.frontend_uri,
    ]
  }
}
