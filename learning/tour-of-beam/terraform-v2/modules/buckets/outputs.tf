output "functions-bucket-id" {
  value = google_storage_bucket.cloud_functions_bucket.id
}

output "functions-bucket-name" {
  value = google_storage_bucket.cloud_functions_bucket.name
}

output "function-bucket-object" {
  value = google_storage_bucket_object.zip.name
}
