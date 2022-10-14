resource "google_cloudfunctions_function" "cloud_function" {
  name                  = "tour-of-beam-backend-cloud-function"
  runtime               = "go116"
  project               = var.project_id
  service_account_email = var.service_account_id
  source_archive_bucket = var.source_archive_bucket
  source_archive_object = var.source_archive_object
  region                = var.region
  ingress_settings      = "ALLOW_ALL"
  # Get the source code of the cloud function as a Zip compression
  trigger_http = true
  # Name of the function that will be executed when the Google Cloud Function is triggered
  entry_point           = "init"
}