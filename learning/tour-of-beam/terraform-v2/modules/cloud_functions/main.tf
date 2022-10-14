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

  entry_point = "getContentTree"
}

# Create IAM entry so all users can invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.cloud_function.project
  region         = google_cloudfunctions_function.cloud_function.region
  cloud_function = google_cloudfunctions_function.cloud_function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}