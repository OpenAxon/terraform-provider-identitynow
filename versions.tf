terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
    }
    identitynow = {
      source = "registry.terraform.fake/axon/identitynow"
      version = "0.1.0"
    }
    vault = {
      source = "hashicorp/vault"
    }
  }
  required_version = ">= 0.13"
}
