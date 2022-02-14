module "lb-http" {
  source  = "GoogleCloudPlatform/lb-http/google//modules/serverless_negs"
  version = "~> 5.1"

  project = var.project_id
  name    = "ms-service"

  ssl                             = var.ssl
  managed_ssl_certificate_domains = [var.domain]
  https_redirect                  = var.ssl
  backends = {
    default = {
      description = null
      groups = [
        {
          group = google_compute_region_network_endpoint_group.serverless_neg.id
        }
      ]
      enable_cdn              = false
      custom_request_headers  = null
      custom_response_headers = null
      security_policy         = null

      iap_config = {
        enable               = false
        oauth2_client_id     = null
        oauth2_client_secret = null
      }

      iap_config = {
        enable               = false
        oauth2_client_id     = null
        oauth2_client_secret = null
      }

      log_config = {
        enable      = false
        sample_rate = null
      }
    }
  }
}

resource "google_compute_region_network_endpoint_group" "serverless_neg" {
  provider              = google-beta
  name                  = "serverless-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "asia-northeast1"
  cloud_run {
    service = google_cloud_run_service.gateway-service.name
  }
}

resource "google_cloud_run_service" "gateway-service" {
  name     = "gateway-service"
  location = var.region
  project  = var.project_id

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal-and-cloud-load-balancing"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/gateway-service"
        env {
          name  = "CATALOG_SERVICE_ADDR"
          value = "catalog-service-y64oiofbkq-an.a.run.app:443"
        }
      }
      service_account_name = "gateway-service@${var.project_id}.iam.gserviceaccount.com"
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}
resource "google_cloud_run_service" "catalog-service" {
  name     = "catalog-service"
  location = var.region
  project  = var.project_id

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/catalog-service"
        env {
          name  = "CUSTOMER_SERVICE_ADDR"
          value = "customer-service-y64oiofbkq-an.a.run.app:443"
        }
        env {
          name  = "ITEM_SERVICE_ADDR"
          value = "item-service-y64oiofbkq-an.a.run.app:443"
        }
      }
      service_account_name = "catalog-service@${var.project_id}.iam.gserviceaccount.com"
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

resource "google_cloud_run_service" "customer-service" {
  name     = "customer-service"
  location = var.region
  project  = var.project_id

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/customer-service"
      }
      service_account_name = "customer-service@${var.project_id}.iam.gserviceaccount.com"
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

resource "google_cloud_run_service" "item-service" {
  name     = "item-service"
  location = var.region
  project  = var.project_id

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/item-service"
        env {
          name  = "DB_NAME"
          value = var.db_name
        }
        env {
          name  = "DB_PASS"
          value = var.db_pass
        }
        env {
          name  = "DB_USER"
          value = var.db_user
        }
        env {
          name  = "INSTANCE_CONNECTION_NAME"
          value = var.instance_connection_name
        }
      }
      service_account_name = "item-service@${var.project_id}.iam.gserviceaccount.com"
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

resource "google_cloud_run_service_iam_member" "public-access" {
  location = google_cloud_run_service.gateway-service.location
  project  = google_cloud_run_service.gateway-service.project
  service  = google_cloud_run_service.gateway-service.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}
resource "google_cloud_run_service_iam_binding" "catalog-service-private-access" {
  location = google_cloud_run_service.catalog-service.location
  project  = google_cloud_run_service.catalog-service.project
  service  = google_cloud_run_service.catalog-service.name
  role     = "roles/run.invoker"
  members = [
    "serviceAccount:gateway-service@${var.project_id}.iam.gserviceaccount.com",
  ]
}
resource "google_cloud_run_service_iam_binding" "customer-service-private-access" {
  location = google_cloud_run_service.customer-service.location
  project  = google_cloud_run_service.customer-service.project
  service  = google_cloud_run_service.customer-service.name
  role     = "roles/run.invoker"
  members = [
    "serviceAccount:catalog-service@${var.project_id}.iam.gserviceaccount.com",
  ]
}
resource "google_cloud_run_service_iam_binding" "item-service-private-access" {
  location = google_cloud_run_service.item-service.location
  project  = google_cloud_run_service.item-service.project
  service  = google_cloud_run_service.item-service.name
  role     = "roles/run.invoker"
  members = [
    "serviceAccount:catalog-service@${var.project_id}.iam.gserviceaccount.com",
  ]
}

resource "google_sql_database" "myinstance-postgres" {
  charset   = "UTF8"
  collation = "en_US.UTF8"
  instance  = google_sql_database_instance.myinstance.name
  name      = "postgres"
  project   = var.project_id
}

resource "google_sql_database_instance" "myinstance" {
  database_version = "POSTGRES_13"
  name             = "myinstance"
  project          = var.project_id
  region           = var.region

  settings {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"

    backup_configuration {
      backup_retention_settings {
        retained_backups = "7"
        retention_unit   = "COUNT"
      }

      binary_log_enabled             = "false"
      enabled                        = "false"
      location                       = "asia"
      point_in_time_recovery_enabled = "false"
      start_time                     = "11:00"
      transaction_log_retention_days = "7"
    }

    database_flags {
      name  = "cloudsql.iam_authentication"
      value = "on"
    }

    disk_autoresize       = "true"
    disk_autoresize_limit = "0"
    disk_size             = "100"
    disk_type             = "PD_SSD"

    ip_configuration {
      authorized_networks {
        name  = "mypgadmin"
        value = "39.110.217.10"
      }

      ipv4_enabled = "true"
      require_ssl  = "false"
    }

    location_preference {
      zone = "asia-northeast1-a"
    }

    pricing_plan = "PER_USE"
    tier         = "db-custom-1-3840"
  }
}
