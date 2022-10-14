#Generates archive of source code
variable "bucket_name" {
  description = "Bucket name to store function source code"
}

variable "location" {
  description = "Cloud Functions Bucket Region"
  default     = "europe-west1"
}

variable "project_id" {
  description = "Our GCP Project"
}

data "archive_file" "source" {
  type        = "zip"
  source_dir  = "../../../backend_"
  output_path = "/tmp/backend.zip"
}