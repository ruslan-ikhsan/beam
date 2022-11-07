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

#Generates archive of source code
variable "bucket_name" {
  description = "The bucket name to store functions' source code"
}

variable "location" {
  description = "Cloud Functions bucket Region"
  default     = "us-central1"
}

variable "project_id" {
  description = "The ID of the Google Cloud project within which resources are provisioned"
}

data "archive_file" "source_code" {
  type        = "zip"
  source_dir  = "../backend"
  output_path = "/tmp/backend.zip"
}