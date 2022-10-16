resource "google_cloudfunctions_function" "cloud_function1" {
  name                  = "tour-of-beam-backend-cloud-function1"
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
  entry_point = "getSdkList"

  environment_variables = {
    DATASTORE_PROJECT_ID="test-cloud-func-deploy"
    TOB_MOCK=1
  }

}

resource "google_cloudfunctions_function" "cloud_function2" {
  name                  = "tour-of-beam-backend-cloud-function2"
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

  environment_variables = {
    DATASTORE_PROJECT_ID="test-cloud-func-deploy"
    TOB_MOCK=1
  }

}

resource "google_cloudfunctions_function" "cloud_function3" {
  name                  = "tour-of-beam-backend-cloud-function3"
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
  entry_point = "getUnitContent"

  environment_variables = {
    DATASTORE_PROJECT_ID="test-cloud-func-deploy"
    TOB_MOCK=1
  }

}

# Create IAM entry so all users can invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker1" {
  project        = google_cloudfunctions_function.cloud_function1.project
  region         = google_cloudfunctions_function.cloud_function1.region
  cloud_function = google_cloudfunctions_function.cloud_function1.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"

  depends_on = [google_cloudfunctions_function.cloud_function1]
}

# Create IAM entry so all users can invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker2" {
  project        = google_cloudfunctions_function.cloud_function2.project
  region         = google_cloudfunctions_function.cloud_function2.region
  cloud_function = google_cloudfunctions_function.cloud_function2.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"

  depends_on = [google_cloudfunctions_function.cloud_function2]
}

# Create IAM entry so all users can invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker3" {
  project        = google_cloudfunctions_function.cloud_function3.project
  region         = google_cloudfunctions_function.cloud_function3.region
  cloud_function = google_cloudfunctions_function.cloud_function3.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"

  depends_on = [google_cloudfunctions_function.cloud_function3]
}