output "service-account-email" {
  value = module.iam.service-account-email
}

output "cloud-function-trigger-url1" {
  value = module.cloud_functions.cloud-function-trigger-url1
}

output "cloud-function-trigger-url2" {
  value = module.cloud_functions.cloud-function-trigger-url2
}

output "cloud-function-trigger-url3" {
  value = module.cloud_functions.cloud-function-trigger-url3
}

output "functions-bucket-name" {
  value = module.buckets.functions-bucket-name
}

output "function-bucket-object" {
  value = module.buckets.function-bucket-object
}