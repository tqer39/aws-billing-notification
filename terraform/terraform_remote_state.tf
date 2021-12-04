data "terraform_remote_state" "management" {
  backend = "remote"
  config = {
    organization = "private-aws"
    workspaces = {
      name = "aws-management"
    }
  }
}