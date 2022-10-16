# [Tour of Beam] Deploying backend infrastructure as GCP Cloud Function with GCP, Terraform and Cloud Build
This tutorial explains how to create Tour of Beam backend infrastructure as a GCP Cloud Functions using Infrastructure as a Code (IaC) and GitOps methodoliges.


### Deployment
Prerequisites:
- Active GCP project with billing enabled
- Existing bucket to store Terraform state (name to be declared in backend.tf)
- Existing service account with the following roles:
    - Cloud Functions Admin
    - Project IAM Admin
    - Service Account Admin
    - Service Account User
    - Storage Admin
    - Storage Object Admin
- Exported JSON Key of service account above

#### Configuring your environment

Steps below will:
1. Enable required APIs for the project
2. Create service account and assign required IAM roles to it (service account to run the function with)
3. Create bucket to archive and store source code
4. Create cloud functions to each function defined in backend source code


```bash
# Create environment directory per your requirements/policy
mkdir environments/dev
cd ../environments/dev
# Import exported JSON Key using
export GOOGLE_APPLICATION_CREDENTIALS = '..key path...'
# Initiliaze and run terraform
terraform init
terraform plan
terraform apply
terraform destroy
```


### Sample usage

Entry point: list sdk names
```
$ curl -X GET https://$REGION-$PROJECT_ID.cloudfunctions.net/getSdkList | json_pp
```
[response](./samples/api/get_sdk_list.json)

Get content tree by sdk name (SDK name == SDK id)
```
$ curl -X GET 'https://$REGION-$PROJECT_ID.cloudfunctions.net/getContentTree?sdk=python'
```
[response](./samples/api/get_content_tree.json)


Get unit content tree by sdk name and unitId
```
$ curl -X GET 'https://$REGION-$PROJECT_ID.cloudfunctions.net/getContentTree?sdk=python&id=challenge1'
```
[response](./samples/api/get_unit_content.json)