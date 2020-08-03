data "vault_generic_secret" "identitynow_client" {
  path = "identitynow/apiclient"
}

provider "vault" {
  # It is strongly recommended to configure this provider through the
  # environment variables described above, so that each user can have
  # separate credentials set in the environment.
  #
  address = "https://10.28.4.12:8200"
  skip_tls_verify = true
}

provider "azurerm" {
  alias = "rm"
  version         = "~> 1.36"
  environment = "usgovernment"
  tenant_id = "5766f879-f51d-4e43-88e6-a8c609103041"
  subscription_id = "4d47ec28-1799-4c7b-9e34-81300079de6b"
  features {}
}

provider "identitynow" {
  api_url = "https://axon-sb.api.identitynow.com"
  client_id = data.vault_generic_secret.identitynow_client.data["client_id"]
  client_secret = data.vault_generic_secret.identitynow_client.data["client_secret"]
}