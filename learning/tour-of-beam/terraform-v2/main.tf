provider "google" {
  project = var.project_id
}

module "iam" {
  source = "modules/iam"
  project_id = var.project_id
  service_account_id = var.service_account_id
  depends_on = [module.api_enable]
}

module "buckets" {
  source = "modules/buckets"
  project_id = var.project_id
  bucket_name = var.bucket_name
  depends_on = [module.iam, module.api_enable]
}

module "api_enable" {
  source = "modules/api_enable"
  project_id = var.project_id
}

module "cloud_functions" {
  source = "modules/cloud_functions"
  region = var.region
  project_id = var.project_id
  bucket_name = var.bucket_name
  service_account_id = module.iam.service-account-email
  source_archive_bucket = module.buckets.functions-bucket-name
  source_archive_object = module.buckets.function-bucket-object
  depends_on = [module.buckets, module.iam, module.api_enable]
}