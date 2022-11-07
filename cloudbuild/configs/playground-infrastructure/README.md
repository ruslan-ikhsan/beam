## Overview

This directory holds Infrastructure-as-Code using terraform to set up cloud build service account and create a trigger


## Requirements

- Existing Google Cloud project with billing enabled
- Existing Cloud Storage Bucket to save Terraform state
- [Google Cloud SDK](https://cloud.google.com/sdk): gcloud init, gcloud auth, and gcloud auth application-default login
- [Terraform](https://www.terraform.io/)


## Usage

Steps below labeled with (manual) are required to be performed once for each target Google Cloud project. Otherwise, the intent is for the module to be executed within an automated process, though can be executed manually as well.

### Commands to run the scripts:

```console
# Create new configuration to auth to GCP Project
foo@bar:~$ gcloud init

```

```console
# Acquire new user credentials to use for Application Default Credentials
foo@bar:~$ gcloud auth application-default login

```

```console
# Setup Cloud Build SA
foo@bar:~$ cd 01.setup
foo@bar:~$ terraform init
foo@bar:~$ terraform plan
foo@bar:~$ terraform apply

```

```console
# Setup Cloud Build Trigger
foo@bar:~$ cd 02.build
foo@bar:~$ terraform init
foo@bar:~$ terraform plan
foo@bar:~$ terraform apply

```


## 01.setup

This directory creates service account that Cloud Build will use in its trigger to deploy Playground Infrastructure.

And assigns necessary IAM roles for that.


## 02.build

This directory creates Cloud Build trigger that will run Cloud Build configuration file to deploy Playground Infrastructure.