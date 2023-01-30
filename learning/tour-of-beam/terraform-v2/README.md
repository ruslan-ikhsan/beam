<!--
    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
-->

# Overview

This project deploys backend infrastructure for Tour of Beam as GCP Cloud Functions using Terraform.

# Requirements

## Development requirements

- [Active GCP project with billing enabled](https://developers.google.com/workspace/guides/create-project)
- [Existing bucket to store Terraform state](https://cloud.google.com/storage/docs/creating-buckets)
- [Existing service account with the following roles](https://cloud.google.com/iam/docs/creating-managing-service-accounts):
    - Cloud Functions Admin
    - Project IAM Admin
    - Service Account Admin
    - Service Account User
    - Storage Admin
    - Storage Object Admin

#### Configuring your environment

Steps below will:
1. Enable required APIs for the project
2. Create service account and assign required IAM roles to it (service account to run the function with)
3. Create bucket to archive and store source code
4. Create cloud functions to each function defined in backend source code


```bash
# Create environment directory per your requirements/policy
mkdir environments/dev
cd ../environments/dev
# Create new configuration to auth to GCP Project
gcloud init
# Acquire new user credentials to use for Application Default Credentials
gcloud auth application-default login
# Initiliaze and run terraform
terraform init
terraform plan
terraform apply
```


### Sample usage

Entry point: list sdk names
```
$ curl -X GET https://$REGION-$PROJECT_ID.cloudfunctions.net/getSdkList | json_pp
```
[response](./samples/api/get_sdk_list.json)

Get content tree by sdk name (SDK name == SDK id)
```
$ curl -X GET 'https://$REGION-$PROJECT_ID.cloudfunctions.net/getContentTree?sdk=python'
```
[response](./samples/api/get_content_tree.json)


Get unit content tree by sdk name and unitId
```
$ curl -X GET 'https://$REGION-$PROJECT_ID.cloudfunctions.net/getContentTree?sdk=python&id=challenge1'
```
[response](./samples/api/get_unit_content.json)