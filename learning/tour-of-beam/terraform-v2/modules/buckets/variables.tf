#Generates archive of source code
variable "name" {
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
  source_dir  = "../../../backend"
  output_path = "/tmp/backend.zip"
}