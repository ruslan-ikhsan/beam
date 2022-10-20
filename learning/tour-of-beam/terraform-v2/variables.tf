variable "bucket_name" {
  description = "The bucket name that will store functions' source code"
}
variable "project_id" {
  type = string
  description = "The ID of the Google Cloud project within which resources are provisioned"
}

variable "iac_service_account_id" {
  type = string
  description = "The ID of the service account responsible for provisioning Google Cloud resources using Infrastructure-as-Code"
  default = "terraform"
}

variable "service_account_id" {
  type = string
  description = "The ID of the service account responsible for running Google Cloud functions"
}
variable "region" {
  default = "us-central1"
}