locals {
  env = "uat"
}

provider "google" {
  project = var.project_id
}

module "iam" {
  source = "../../modules/iam"
  project_id = var.project_id
  service_account_id = var.service_account
  depends_on = [module.buckets]
}

module "buckets" {
  source = "../../modules/buckets"
  project_id = var.project_id
}

module "api_enable" {
  source = "../../modules/api_enable"
  project_id = var.project_id
  depends_on = [module.buckets, module.iam]
}

module "cloud_functions" {
  source = "../../modules/cloud_functions"
  depends_on = [module.buckets, module.iam, module.api_enable]
}