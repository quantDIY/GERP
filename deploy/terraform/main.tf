# deploy/terraform/main.tf
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

variable "project_id" {
  type    = string
  default = "gerp-production"
}

variable "region" {
  type    = string
  default = "us-central1"
}

# Provision the Production Spanner Matrix
resource "google_spanner_instance" "gerp_prod_instance" {
  config       = "regional-${var.region}"
  display_name = "GERP Production Instance"
  num_nodes    = 3
  name         = "gerp-prod-instance"
  labels = {
    environment = "production"
    domain      = "gerp-matrix"
  }
}

resource "google_spanner_database" "gerp_prod_db" {
  instance = google_spanner_instance.gerp_prod_instance.name
  name     = "gerp-prod-db"
  
  # Note: DDL structurally managed securely via the GERP physical CLI migrations. 
  # Handled mathematically out-of-band to prevent strict state synchronization failures directly in Terraform.
}
