terraform {
  backend "remote" {
    organization = "example-org-7cd94b"

    workspaces {
      name = "microservices-pattern"
    }
  }
}

