variable "region" {}
variable "project_id" {}
variable "service_account_id" {
  description = "Name of SA to run Cloud Function"
}
variable "bucket_name" {}
variable "source_archive_bucket" {}

variable "source_archive_object" {}