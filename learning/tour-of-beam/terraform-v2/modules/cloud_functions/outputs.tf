output "cloud-function-trigger-url1" {
  value = google_cloudfunctions_function.cloud_function1.https_trigger_url
}

output "cloud-function-trigger-url2" {
  value = google_cloudfunctions_function.cloud_function2.https_trigger_url
}

output "cloud-function-trigger-url3" {
  value = google_cloudfunctions_function.cloud_function3.https_trigger_url
}