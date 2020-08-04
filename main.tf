data "vault_generic_secret" "aad_client" {
  path = "identitynow/aadclient"
}

resource "identitynow_source_azure_ad" "source_azure_ad_ag1" {
  name = "Azure AD product (usgov) TERRAFORM TEST 8"
  description = "The Azure Active Directory connector for the tenant with all US gov subscriptions on the product side. TERRAFORM TEST 8a"

  owner {
    id = "2c91808472ed35250172f1e9ec947b22"
    name = "greg_burton"
    type = "IDENTITY"
  }

  cluster {
    id = "2c91808672dd6f5d017308a1d2d25d35"
    name = "sbx-product-ag1"
  }

  connector_attributes {
    grant_type = "CLIENT_CREDENTIALS"
    client_id = data.vault_generic_secret.aad_client.data["client_id"]
    client_secret = data.vault_generic_secret.aad_client.data["client_secret"]
    domain_name = "us1.axonengineering.io"
    ms_graph_resource_base = "https://graph.microsoft.us"
    ms_graph_token_base = "https://login.microsoftonline.us"
    azure_ad_graph_resource_base = "https://graph.microsoftazure.us"
    azure_ad_graph_token_base = "https://login.microsoftonline.us"
  }
}
