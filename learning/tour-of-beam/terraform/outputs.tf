#
# Licensed to the Atpache Software Foundation (ASF) under one
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
#

output "cloud_function_trigger_url_1" {
  value = google_cloudfunctions_function.backend_function_1.https_trigger_url
}

output "cloud_function_trigger_url_2" {
  value = google_cloudfunctions_function.backend_function_2.https_trigger_url
}

output "cloud_function_trigger_url_3" {
  value = google_cloudfunctions_function.backend_function_3.https_trigger_url
}