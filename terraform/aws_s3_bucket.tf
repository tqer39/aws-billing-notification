resource "aws_s3_bucket" "sources" {
  provider      = aws.billing-notification
  bucket        = "sources-${data.terraform_remote_state.management.outputs.organizations_account_management_id}"
  acl           = "private"
  force_destroy = true
  versioning {
    enabled = true
  }
}