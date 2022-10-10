# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

provider "google" {
  project = var.project_id
  region = var.region
}

resource "google_cloudfunctions_function" "backend_function_1" {
  name                  = "getSdkList-function"
  runtime               = "go116"
  region                = var.region
  service_account_email = var.service_account_email
  ingress_settings      = "ALLOW_ALL"
  # Get the source code of the cloud function as a Zip compression
  source_archive_bucket = google_storage_bucket.function_bucket.name
  source_archive_object = google_storage_bucket_object.zip.name

  trigger_http = true
  # Name of the function that will be executed when the Google Cloud Function is triggered (def hello_gcs)
  entry_point           = "getSdkList"

}

resource "google_cloudfunctions_function" "backend_function_2" {
  name                  = "getContentTree-function"
  runtime               = "go116"
  region                = var.region
  service_account_email = var.service_account_email
  ingress_settings      = "ALLOW_ALL"
  # Get the source code of the cloud function as a Zip compression
  source_archive_bucket = google_storage_bucket.function_bucket.name
  source_archive_object = google_storage_bucket_object.zip.name

  trigger_http = true
  # Name of the function that will be executed when the Google Cloud Function is triggered (def hello_gcs)
  entry_point           = "getContentTree"

}

resource "google_cloudfunctions_function" "backend_function_3" {
  name                  = "getUnitContent-function"
  runtime               = "go116"
  region                = var.region
  service_account_email = var.service_account_email
  ingress_settings      = "ALLOW_ALL"
  # Get the source code of the cloud function as a Zip compression
  source_archive_bucket = google_storage_bucket.function_bucket.name
  source_archive_object = google_storage_bucket_object.zip.name

  trigger_http = true
  # Name of the function that will be executed when the Google Cloud Function is triggered (def hello_gcs)
  entry_point           = "getUnitContent"

}
