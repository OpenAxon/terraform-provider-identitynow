terraform {
  required_version = "~> 0.12"
  backend "azurerm" {
    environment = "usgovernment"
    resource_group_name  = "qag1ge1-Storage-Terraform"
    storage_account_name = "qag1ge1terraform"
    container_name = "tfstate-identitynow"
    key            = "terraform.tfstate"
  }
}

data "azurerm_resource_group" "terraform_storage" {
  name = format("%s%s%s-Storage-Terraform", upper(var.environment_code), upper(var.deployment_code), upper(var.location_code))
}

data "azurerm_storage_account" "terraform_storage_account" {
  name                = format("%s%s%sterraform", var.environment_code, var.deployment_code, var.location_code)
  resource_group_name = data.azurerm_resource_group.terraform_storage.name
}
