resource "identitynow_source_azure_ad" "source_azure_ad_ag1" {
  name = "Azure AD product (usgov) TERRAFORM TEST"
  description = "The Azure Active Directory connector for the tenant with all US gov subscriptions on the product side. TERRAFORM TEST"
  owner {
    name = "greg_burton"
  }
}
