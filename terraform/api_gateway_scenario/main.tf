variable "REGION" { default="us-east-1" }
variable "PROFILE" { default="default" }

provider "aws" {
  profile = var.PROFILE
  region = var.REGION
}
