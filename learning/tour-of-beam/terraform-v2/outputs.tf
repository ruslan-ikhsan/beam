output "service-account-email" {
  value = module.iam.service-account-email
}

output "cloud-function-trigger-url" {
  value = module.cloud_functions.cloud-function-trigger-url
}

output "functions-bucket-name" {
  value = module.buckets.functions-bucket-name
}

output "function-bucket-object" {
  value = module.buckets.function-bucket-object
}