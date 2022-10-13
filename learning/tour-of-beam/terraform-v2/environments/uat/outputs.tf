output "service-account-email" {
  value = module.iam.service-account-email
}

output "tf-state-bucket-id" {
  value = module.buckets.tf-state-bucket-id
}

output "tf-state-bucket-name" {
  value = module.buckets.tf-state-bucket-name
}