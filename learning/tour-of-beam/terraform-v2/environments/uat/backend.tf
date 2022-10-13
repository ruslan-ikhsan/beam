terraform {
  backend "gcs" {
    bucket = var.bucket_name
    prefix = "env/uat"
  }
}