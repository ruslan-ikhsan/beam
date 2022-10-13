resource "google_storage_bucket" "bucket_for_function" {
  name          = var.bucket_name
  location      = var.location
  project       = var.project_id
  storage_class = "STANDARD"
  force_destroy = true
}

resource "google_storage_bucket_object" "zip" {
  # Use an MD5 here. If there's no changes to the source code, this won't change either.
  # We can avoid unnecessary redeployments by validating the code is unchanged, and forcing
  # a redeployment when it has!
  name         = "${data.archive_file.source.output_md5}.zip"
  bucket       = google_storage_bucket.bucket_for_function.name
  source       = data.archive_file.source.output_path
  content_type = "application/zip"
}

resource "google_storage_bucket_access_control" "public_rule" {
  bucket = google_storage_bucket.bucket_for_function.name
  role   = "READER"
  entity = "allUsers"
}