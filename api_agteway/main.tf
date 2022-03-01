variable "region" { default="us-east-1" }
variable "profile" { default="default" }

provider "aws" {
  profile = var.profile
  region = var.region
}
