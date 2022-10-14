resource "google_cloudfunctions_function" "cloud_function" {
  name                  = "tour-of-beam-backend-cloud-function"
  runtime               = "go116"
  project               = var.project_id
  region                = var.region
  service_account_email = module.iam.service-account-email
  ingress_settings      = "ALLOW_ALL"
  # Get the source code of the cloud function as a Zip compression
  source_archive_bucket = module.buckets.functions-bucket-name
  source_archive_object = module.buckets.function-bucket-object

  trigger_http = true
  # Name of the function that will be executed when the Google Cloud Function is triggered
  entry_point           = "init"

}

module "iam" {
  source = "../iam"
  service_account_id = var.service_account_id
  project_id = var.project_id
}

module "buckets" {
  source = "../buckets"
  project_id = var.project_id
  bucket_name = module.buckets.functions-bucket-name
}