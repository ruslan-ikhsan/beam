variable "project_id" {
  description = "Our GCP Project"
}

variable "service_account_id" {
  description = "Name of SA to run Cloud Function"
  default = "tourofbeam-cfsa"
}

