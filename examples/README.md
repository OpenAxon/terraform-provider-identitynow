Example command to run plan and apply for resources:

`terraform plan -var="api_client_id=<CLIENT_ID>" -var="api_client_secret=<CLIENT_SECRET>"`

`terraform apply -var="api_client_id=<CLIENT_ID>" -var="api_client_secret=<CLIENT_SECRET>"`

To list resources:

```
$ terraform state list
data.identitynow_identity.john_doe
data.identitynow_source_entitlement.aad_operator
data.identitynow_source_entitlement.ad_developer
identitynow_access_profile.ad_access_profile_developers
identitynow_access_profile.aad_access_profile_operators
identitynow_account_aggregation_schedule.account_aggregation
identitynow_account_schema_attribute.account_schema_employeeId_attribute
identitynow_password_policy.password_policy
identitynow_role.operator_developer_role
identitynow_source.active_directory_source
identitynow_source.aws_iam_source
identitynow_source.azure_ad_source
```
