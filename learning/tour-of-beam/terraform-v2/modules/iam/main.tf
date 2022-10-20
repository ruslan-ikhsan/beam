resource "google_service_account" "cloud_function_sa" {
  account_id   = var.service_account_id
  display_name = "Service Account to run Cloud Functions"
}

resource "google_project_iam_member" "terraform_service_account_roles" {
  for_each = toset([
    "roles/cloudfunctions.admin", "roles/storage.objectViewer",
    "roles/storage.objectCreator", "roles/iam.serviceAccountUser"
  ])
  role    = each.key
  member  = "serviceAccount:${google_service_account.cloud_function_sa.email}"
  project = var.project_id
}

