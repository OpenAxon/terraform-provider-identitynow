# Usage
This POC is simple, so it's hardwired to AG1 environment right now. Upon moving to a second environment, it can be enhanced to be configurable.

This Terraform configuration for IdentityNow stores its state remotely in Azure blob storage. Most configuration is passed directly to the backend, but there was one value ("environment")
that does not work for some reason. The equivalent env variable does work, so it needs to be set to run this Terraform code.
```bash
export ARM_ENVIRONMENT=usgovernment
```

Simple commands to manage resources:
```bash
terraform plan
terraform apply
terraform state list
terraform destroy
```

Example declaration of an Azure Active Directory source connector:
```hcl-terraform
data "vault_generic_secret" "aad_client" {
  path = "identitynow/aadclient"
}

resource "identitynow_source" "source_azure_ad_ag1" {
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
```

# Development
Edit the Go files that make up the provider, and rebuild/reload the provider:
```bash
go build -o terraform-provider-identitynow
terraform init
``` 

### Debugging
To debug the provider logic, set `TF_LOG=TRACE`.

# Set up secrets in Vault
Execute the following while logged into Vault with the all-powerful root token. We want to do minimal steps with the root token, then create a new, limited token that we will use from Terraform.
```bash
./tools.sh
export VAULT_ADDR=https://qag1ge1lvlt001.ag1.taservs.net:8200
export VAULT_SKIP_VERIFY=true
vault login

vault secrets enable -path=identitynow kv
vault kv put identitynow/apiclient client_id=c70cde50e14d4e5e9082392056f9faf3 client_secret=<REDACTED>
vault kv put identitynow/aadclient client_id=6e732858-b263-4c9e-b752-a229626e18a7 client_secret=<redacted>
```

### Create an appropriate policy for IdentityNow Terraform provider
```bash
cat <<HERE | vault policy write identitynow -
<paste policy from vault/identitynow_policy.hcl>
HERE
```

### Create a token with access to only to IdentityNow secrets
```bash
vault token create -policy=identitynow
Key                  Value
---                  -----
token                <REDACTED>
token_accessor       mBBuqHL5vYpwvgclwnXR8vID
token_duration       768h
token_renewable      true
token_policies       ["default" "identitynow"]
identity_policies    []
policies             ["default" "identitynow"]
```
View the token you just created:
```bash
vault token lookup -accessor mBBuqHL5vYpwvgclwnXR8vID
```
Place the sensitive token value into ~/.vault-token, so that it will be used by Terraform and the Vault CLI for authenticating. This type of token will eventually be linked to LDAP auth of each user. For now, we'll just pass this one around.





