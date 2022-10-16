resource "google_cloudfunctions_function" "cloud_function" {
  count                 = length(var.entry_point_names)
  name                  = "tour-of-beam-backend-cloud-function-${count.index}"
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
  entry_point = var.entry_point_names[count.index]

  environment_variables = {
    DATASTORE_PROJECT_ID="test-cloud-func-deploy"
    TOB_MOCK=1
  }

}

# Create IAM entry so all users can invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker" {
  count          = length(google_cloudfunctions_function.cloud_function)
  project        = google_cloudfunctions_function.cloud_function[count.index].project
  region         = google_cloudfunctions_function.cloud_function[count.index].region
  cloud_function = google_cloudfunctions_function.cloud_function[count.index].name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"

  depends_on = [google_cloudfunctions_function.cloud_function]
}