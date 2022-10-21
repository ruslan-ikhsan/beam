resource "google_storage_bucket" "cloud_functions_bucket" {
  name          = var.bucket_name
  location      = var.location
  project       = var.project_id
  storage_class = "STANDARD"
}

resource "google_storage_bucket_object" "zip" {
  # Use an MD5 here. If there's no changes to the source code, this won't change either.
  # We can avoid unnecessary redeployments by validating the code is unchanged, and forcing
  # a redeployment when it has!
  name         = "${data.archive_file.source_code.output_md5}.zip"
  bucket       = google_storage_bucket.cloud_functions_bucket.name
  source       = data.archive_file.source_code.output_path
  content_type = "application/zip"
}