variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-east-1"  
}

variable "aws_ecr_repository_name" {
  description = "The name of the ecr repository"
  type        = string
  default     = "httpserver"  
}

variable "dkr_src" {
  description = "The directory of Docker file and source code"
  type        = string
  default     = "/Users/joseph/go-projects/Tojo_Challenge/Infrastructure/http_server"  
}

variable "app_count" {
  type = number
  default = 1
}

