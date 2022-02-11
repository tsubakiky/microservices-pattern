terraform {
  backend "remote" {
    organization = "example-org-7cd94b"

    workspaces {
      name = "microservices-pattern"
    }
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.53"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 3.53"
    }
  }
  required_version = ">= 0.13"
}
