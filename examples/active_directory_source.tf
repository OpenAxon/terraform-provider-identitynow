resource "identitynow_source" "active_directory_source" {
  name        = "Active Directory Source"
  description = "The Active Directory connector created by terraform"
  connector   = "active-directory"

  owner {
    id   = data.identitynow_identity.john_doe.external_id
    name = data.identitynow_identity.john_doe.name
    type = "IDENTITY"
  }

  cluster {
    id   = "<CLUSTER_ID>"
    name = "<CLUSTER_NAME>"
  }

  password_policies {
    id   = identitynow_password_policy.password-policy.id
    name = identitynow_password_policy.password-policy.name
    type = "PASSWORD_POLICY"
  }

  connector_attributes {
    iq_service_host        = "iqservice.example.com"
    iq_service_port        = "5052"
    iq_service_user        = "<IQ_SERVICE_USER>"
    iq_service_password    = "<IQ_SERVICE_PASSWORD>"
    use_tls_for_iq_service = true
    encrypted              = "IQServicePassword"
    forest_settings {
      user               = "<IQ_SERVICE_USER>"
      password           = "<IQ_SERVICE_PASSWORD>"
      gc_server          = "ad1.example.com:3269"
      forest_name        = "example.com"
      use_ssl            = true
      authorization_type = "simple"
    }
    forest_settings {
      user               = "<IQ_SERVICE_USER>"
      password           = "<IQ_SERVICE_PASSWORD>"
      gc_server          = "ad2.example.com:3269"
      forest_name        = "example.com"
      use_ssl            = true
      authorization_type = "simple"
    }
    domain_settings {
      user     = "<IQ_SERVICE_USER>"
      password = "<IQ_SERVICE_PASSWORD>"
      servers = [
        "ad1.example.com:3269",
        "ad2.example.com:3269"]
      port               = "636"
      forest_name        = "example.com"
      authorization_type = "simple"
      domain_dn          = "DC=example,DC=com"
      use_ssl            = true
    }
    search_dns {
      search_dn             = "OU=Users,DC=example,DC=com"
      iterate_search_filter = "(objectclass=person)"
      search_scope          = "SUBTREE"
    }
    search_dns {
      search_dn             = "OU=OffboardedUsers,DC=example,DC=com"
      iterate_search_filter = "(objectclass=person)"
      search_scope          = "SUBTREE"
    }
    group_search_dns {
      search_dn = "OU=Groups,DC=example,DC=com"
      iterate_search_filter = "(objectclass=group)"
      search_scope          = "SUBTREE"
    }
  }

  lifecycle {
    ignore_changes = [
      connector_attributes[0].iq_service_password,
      connector_attributes[0].forest_settings[0].password,
      connector_attributes[0].forest_settings[1].password,
      connector_attributes[0].domain_settings[0].password,
      management_workgroup[0].name,
    ]
  }
}
