variable "bucket_name" {
  description = "Bucket name to store function source code"
}
variable "project_id" {
  description = "Our GCP Project"
}

variable "service_account" {
  description = "Name of SA to run Cloud Function"
  default = "newnameforsa"
}

variable "region" {
  default = "europe-west1"
}