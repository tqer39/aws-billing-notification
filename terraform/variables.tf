variable "Environment" {
  description = "環境（dev | stg | prd）の種類"
  type        = string
  default     = "prd"
}

variable "IaC" {
  description = "IaC（Infrastructure as Code）で使用したソフトウェア"
  type        = string
  default     = "Terraform"
}

variable "TFC_Organization" {
  description = "Terraform Cloud >> Organization の名前"
  type        = string
  default     = "private-aws"
}

variable "TFC_Workspace" {
  description = "Terraform Cloud >> Workspace の名前"
  type        = string
  default     = "aws-billing-notification"
}

variable "region" {
  description = "デフォルトのリージョン"
  default     = "ap-northeast-1"
  type        = string
}