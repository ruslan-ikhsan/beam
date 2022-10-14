# Overview

This directory holds Infrastructure-as-Code using terraform to set up, build,
and deploy the application and resource dependencies.

# Requirements

- Existing Google Cloud project with billing enabled
- [Google Cloud SDK](https://cloud.google.com/sdk); `gcloud init`,
`gcloud auth`, and `gcloud auth application-default login`
- [terraform](https://www.terraform.io/)

# Usage

Folders in this [infrastructure](.) directory are ordered by required execution.
Steps below labeled with (manual) are required to be performed once for each
target Google Cloud project.  Otherwise, the intent is for the module to be
executed within an automated process, though can be executed manually as well.

## 01.Setup


