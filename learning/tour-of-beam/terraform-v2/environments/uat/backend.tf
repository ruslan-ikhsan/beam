terraform {
  backend "gcs" {
    bucket = "tour-of-beam-backend-tfstate-bucket"
    prefix = "env/uat"
  }
}