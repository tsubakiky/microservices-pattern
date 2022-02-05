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
      version = "4.9.0"
    }
  }
}
