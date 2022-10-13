output "functions-bucket-id" {
  value = google_storage_bucket.bucket_for_function.id
}

output "functions-bucket-name" {
  value = google_storage_bucket.bucket_for_function.name
}

output "function-bucket-object" {
  value = google_storage_bucket_object.zip.name
}
