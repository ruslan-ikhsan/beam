#Generates archive of source code
variable "bucket_name" {
  description = "The bucket name to store functions' source code"
}

variable "location" {
  description = "Cloud Functions bucket Region"
  default     = "us-central1"
}

variable "project_id" {
  description = "The ID of the Google Cloud project within which resources are provisioned"
}

data "archive_file" "source_code" {
  type        = "zip"
  source_dir  = "../backend"
  output_path = "/tmp/backend.zip"
}