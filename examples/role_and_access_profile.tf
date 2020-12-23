resource "identitynow_access_profile" "ad_access_profile_developers" {
  name = "Active Directory Developer Access Profile"
  description = "Developer Access Profile Description"
  entitlements = [
    data.identitynow_source_entitlement.ad_developer.id]
  source_id = tonumber(identitynow_source.active_directory_source[0].connector_attributes[0].cloud_external_id)
  owner_id  = data.identitynow_identity.john_doe.id
}

resource "identitynow_access_profile" "aad_access_profile_operators" {
  name = "Azure Active Directory Operator Access Profile"
  description = "Operator Access Profile Description"
  entitlements = [
    data.identitynow_source_entitlement.aad_operator.id]
  source_id = tonumber(identitynow_source.azure_ad_source[0].connector_attributes[0].cloud_external_id)
  owner_id  = data.identitynow_identity.john_doe.id
}

resource "identitynow_role" "operator_developer_role" {
  access_profile_ids = [
    identitynow_access_profile.aad_access_profile_operators.id,
    identitynow_access_profile.ad_access_profile_developers[count.index].id
  ]
  description      = "Developer Operator Role Description"
  name             = "Developer Operator Role"
  approval_schemes = "none"
  disabled         = false
  requestable      = true
  owner            = data.identitynow_identity.john_doe.alias
  lifecycle {
    ignore_changes = [
      name,
      display_name,
      identity_count
    ]
  }
}
