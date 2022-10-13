locals {
  env = "uat"
}

provider "google" {
  project = var.project_id
}

module "iam" {
  source = "../../modules/iam"
}

module "buckets" {
  source = "../../modules/buckets"
}

#module "api_enable" {
#  source = "../../modules/api_enable"
#}
