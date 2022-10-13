#Generates archive of source code
variable "name" {
  description = "Bucket name to store TF state"
  default     = "tour-of-beam-backend-tfstate-bucket"
}

variable "location" {
  description = "Cloud Functions Bucket Region"
  default     = "europe-west1"
}

variable "project_id" {
  description = "Our GCP Project"
}