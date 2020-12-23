resource "identitynow_source" "azure_ad_source" {
  name        = "Azure Active Directory Source"
  description = "The Azure Active Directory connector created by terraform"
  connector   = "azure-active-directory"

  owner {
    id   = data.identitynow_identity.john_doe.external_id
    name = data.identitynow_identity.john_doe.name
    type = "IDENTITY"
  }

  cluster {
    id   = "<CLUSTER_ID>"
    name = "<CLUSTER_NAME>"
  }

  #only if there is any management group defined for source
  management_workgroup {
    id   = "<MANAGEMENT_GROUP_ID>"
    name = "<MANAGEMENT_GROUP_NAME>"
    type = "GOVERNANCE_GROUP"
  }

  password_policies {
    id   = identitynow_password_policy.password-policy.id
    name = identitynow_password_policy.password-policy.name
    type = "PASSWORD_POLICY"
  }

  connector_attributes {
    grant_type                   = "CLIENT_CREDENTIALS"
    client_id                    = "<CLIENT_ID>"
    client_secret                = "<CLIENT_SECRET"
    domain_name                  = "example.come"
    ms_graph_resource_base       = "https://graph.microsoft.us" #for commercial https://graph.microsoft.com
    ms_graph_token_base          = "https://login.microsoftonline.us" #for commercial https://login.windows.net
    azure_ad_graph_resource_base = "https://graph.microsoftazure.us" #for commercial https://graph.windows.net
    azure_ad_graph_token_base    = "https://login.microsoftonline.us" #for commercial https://login.microsoftonline.com
    api_version                  = "1.6"
    encrypted                    = "clientSecret" #more fields can be add here separated with `,`
  }
  lifecycle {
    ignore_changes = [
      connector_attributes[0].client_secret,
      connector_attributes[0].encrypted,
      management_workgroup[0].name,
    ]
  }
}

resource "identitynow_password_policy" "password_policy" {
  name                    = "Primary Password Policy"
  description             = "Password Policy for Azure Active Directory and Active Directory Sources"
  min_alpha               = 1
  min_length              = 14
  min_lower               = 1
  min_numeric             = 1
  min_special             = 1
  min_upper               = 1
  use_account_attributes  = true #if true it prevents the use of account attributes
  use_identity_attributes = true #if true it prevents the use of identity attributes
  use_history             = 12
}

#to add a new attribute to source account schema, this resource only can be apply after the source is created and tested connection is working
resource "identitynow_account_schema_attribute" "account_schema_employeeId_attribute" {
  source_id   = identitynow_source.azure_ad_source[0].connector_attributes[0].cloud_external_id
  name        = "employeeId"
  description = "Employee ID"
  type        = "string"
  object_type = "account"
}

#to schedule an account aggregation
resource "identitynow_account_aggregation_schedule" "account_aggregation" {
  source_id        = identitynow_source.azure_ad_source[0].connector_attributes[0].cloud_external_id
  cron_expressions = ["0 0 8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,0,1,2,3,4,5,6,7 * * ?"] #schedule aggregation every hour
}
