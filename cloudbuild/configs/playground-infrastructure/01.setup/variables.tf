variable "project" {
  type = string
  description = "The ID of the Google Cloud project within which resources are provisioned"
}
variable "iac_service_account_id" {
  type = string
  description = "The ID of the service account responsible for provisioning Google Cloud resources using Infrastructure-as-Code"
  default = "terraform"
}