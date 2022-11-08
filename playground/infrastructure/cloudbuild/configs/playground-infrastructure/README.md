
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
# Beam Playground CloudBuild configuration

Beam Playground uses [Google Cloud Build](https://cloud.google.com/build) for application build and deployment. This Readme explains how to connect a GitHub repository with Cloud Build and configure Cloud Build.


## Prerequisites:

- [GCP project](https://cloud.google.com/)
- Enabled [Cloud Build API](https://cloud.google.com/apis/docs/getting-started#enabling_apis)
- [gcloud CLI](https://cloud.google.com/sdk/docs/install-sdk)
- An existing Google Cloud Storage Bucket to save Terraform state - `state_bucket`
- [Terraform](https://www.terraform.io/)
- If you are initially connecting your GitHub repository to Cloud Build, make sure you have admin-level permissions on your repository. To learn more about GitHub repository permissions, see [Repository permission levels for an organization](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/repository-permission-levels-for-an-organization#permission-levels-for-repositories-owned-by-an-organization)
.

## Usage

### Connect Cloud Build and GitHub repository

Please follow these steps to configure CloudBuild for Beam Playground - [Connect to a GitHub repository](https://cloud.google.com/build/docs/automating-builds/github/connect-repo-github)

### Prepare gcloud CLI

```console
# Create new configuration to auth to GCP Project
foo@bar:~$ gcloud init

```

```console
# Acquire new user credentials to use for Application Default Credentials
foo@bar:~$ gcloud auth application-default login

```

### Perform One-time setup steps

1. Navigate to `01.setup` folder
```console
# Navigate to 01.setup folder
foo@bar:~$ cd 01.setup
```
2. Create and setup Google Cloud service account that Cloud Build trigger will use
```console
# Setup Cloud Build SA
foo@bar:~$ terraform init
foo@bar:~$ terraform plan
foo@bar:~$ terraform apply

```
3. Navigate to `02.build` folder
```console
# Navigate to 02.build folder
foo@bar:~$ cd ../02.build
```
4. Create and setup Google Cloud Build trigger that will run Cloud Build configuration file to deploy Playground Infrastructure.
```console
# Setup Cloud Build Trigger
foo@bar:~$ terraform init
foo@bar:~$ terraform plan
foo@bar:~$ terraform apply

```

After completing these steps, GCP project will have Cloud Build trigger that can be used to build and deploy Beam Playground.
