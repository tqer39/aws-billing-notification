terraform {
  required_version = "~> 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3"
    }
  }
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "private-aws"
    workspaces {
      name = "aws-billing-notification"
    }
  }
}

provider "aws" {
  region = var.region
  default_tags {
    tags = local.common_tags
  }
}

provider "aws" {
  alias  = "billing-notification"
  region = var.region
  assume_role {
    role_arn = "arn:aws:iam::${data.terraform_remote_state.management.outputs.organizations_account_management_id}:role/terraform"
  }
  default_tags {
    tags = local.common_tags
  }
}

locals {
  common_tags = {
    IaC              = var.IaC
    Environment      = var.Environment
    TFC_Organization = var.TFC_Organization
    TFC_Workspace    = var.TFC_Workspace
  }
}