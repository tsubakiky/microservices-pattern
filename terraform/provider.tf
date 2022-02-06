provider "google" {
  project     = var.project_id
  credentials = var.GOOGLE_CREDENTIALS
}

provider "google-beta" {
  project     = var.project_id
  credentials = var.GOOGLE_CREDENTIALS
}
