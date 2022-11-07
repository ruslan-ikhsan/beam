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

variable "bucket_name" {
  description = "The bucket name that will store functions' source code"
}

variable "project_id" {
  type = string
  description = "The ID of the Google Cloud project within which resources are provisioned"
}

variable "service_account_id" {
  type = string
  description = "The ID of the service account responsible for running Google Cloud functions"
}

variable "impersonated_account_id" {
  description = "The ID of the impersonated service account that will run TF scripts to deploy infrastructure"
}

variable "region" {
  default = "us-central1"
}

data "google_service_account_access_token" "default" {
  provider               	= google.impersonation
  target_service_account 	= local.terraform_service_account
  scopes                 	= ["userinfo-email", "cloud-platform"]
  lifetime               	= "1200s"
}