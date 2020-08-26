data "vault_generic_secret" "aad_client" {
  path = "identitynow/aadclient"
}

data "vault_generic_secret" "iqservice_client" {
  path = "identitynow/iqserviceclient"
}

resource "identitynow_source" "source_azure_ad_ag1" {
  name = "Product platform, Azure portal, US Gov"
  description = "The Azure Active Directory connector created by terraform for the tenant with all US gov subscriptions on the product side."
  connector = "azure-active-directory"

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
resource "identitynow_source" "source_on-prem_ad_ag1" {
  name = "Product platform, Active Directory, AG1"
  description = "The on-prem Active Directory connector created by terraform for the tenant with all US gov subscriptions on the product side."
  connector = "active-directory"

  owner {
    id = "2c9180847351503d0173542cf3a602b3"
    name = "Elham Amouhadi"
    type = "IDENTITY"
  }

  cluster {
    id = "2c91808672dd6f5d017308a1d2d25d35"
    name = "sbx-product-ag1"
  }

  connector_attributes {
    iq_service_host = "qag1ge1widn001.ag1.taservs.net"
    iq_service_port = "5052"
    iq_service_user = data.vault_generic_secret.iqservice_client.data["iq_service_user"]
    iq_service_password = data.vault_generic_secret.iqservice_client.data["iq_service_password"]
    use_tls_for_iq_service = true
    forest_settings {
        user = data.vault_generic_secret.iqservice_client.data["iq_service_user"]
        password = data.vault_generic_secret.iqservice_client.data["iq_service_password"]
        gc_server = "qag1ge1wact001.ag1.taservs.net:3269"
        forest_name = "ag1.taservs.net"
        use_ssl = true
        authorization_type = "simple"
    }
    forest_settings {
        user = data.vault_generic_secret.iqservice_client.data["iq_service_user"]
        password = data.vault_generic_secret.iqservice_client.data["iq_service_password"]
        gc_server = "qag1ge1wact002.ag1.taservs.net:3269"
        forest_name = "ag1.taservs.net"
        use_ssl = true
        authorization_type = "simple"
    }
    domain_settings {
      user = data.vault_generic_secret.iqservice_client.data["iq_service_user"]
      password = data.vault_generic_secret.iqservice_client.data["iq_service_password"]
      servers= ["qag1ge1wact001.ag1.taservs.net", "qag1ge1wact002.ag1.taservs.net"]
      port = "636"
      forest_name = "ag1.taservs.net"
      authorization_type = "simple"
      domain_dn = "DC=ag1,DC=taservs,DC=net"
      use_ssl = true
    }
    search_dns {
      search_dn = "OU=Users,OU=Evidence.com,DC=ag1,DC=taservs,DC=net"
      iterate_search_filter = "(objectclass=person)"
    }
  }
}
