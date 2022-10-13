variable "region" {}
variable "project_id" {}
variable "bucket_name" {
  description = "Bucket name to store function source code"
}

variable "service_account_id" {
  description = "Name of SA to run Cloud Function"
}