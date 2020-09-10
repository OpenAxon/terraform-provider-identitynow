Example declaration of an Azure Active Directory source connector:
```hcl-terraform
data "vault_generic_secret" "aad_client" {
  path = "identitynow/aadclient"
}

resource "identitynow_source" "source_azure_ad" {
  name = "Azure AD product (usgov)"
  description = "The Azure Active Directory connector for the tenant with all US gov subscriptions on the product side."

  owner {
    id = "2c91808472ed12345672f1e9ec947b22"
    name = "john_doe"
    type = "IDENTITY"
  }

  cluster {
    id = "2c91808672dd1234567308a1d2d25d35"
    name = "product-qa"
  }

  connector_attributes {
    grant_type = "CLIENT_CREDENTIALS"
    client_id = data.vault_generic_secret.aad_client.data["client_id"]
    client_secret = data.vault_generic_secret.aad_client.data["client_secret"]
    domain_name = "us1.example.io"
    ms_graph_resource_base = "https://graph.microsoft.us"
    ms_graph_token_base = "https://login.microsoftonline.us"
    azure_ad_graph_resource_base = "https://graph.microsoftazure.us"
    azure_ad_graph_token_base = "https://login.microsoftonline.us"
  }
}
```

# Development
Edit the Go files that make up the provider, and rebuild/reload the provider.
```bash
./build.sh
```
 
In Terraform 13, the convenience of looking in the current working directory for the provider binary is no longer supported. Instead, you must build the binary and place it in a specific path structure:
```
~/.terraform.d/plugins/registry.terraform.fake/axon/identitynow/0.1.0/darwin_amd64/terraform-provider-identitynow_v0.1.0
```
The path structure format adheres to strict conventions. When we reach the point of publishing the provider to registry.terraform.io, we can execute the following to update the provider on all resources in TF state:
```
terraform state replace-provider "registry.terraform.fake/axon/identitynow" "registry.terraform.io/axon/identitynow"
```

# Testing the Provider

In order to test the provider, you can simply run `make test`.
```sh
$ make test
```
In order to run the full suite of Acceptance test identitynow url, client id and secret, owner name and id, cluster name and id are needed to make the API call to create IdentityNow source for test.

To run acceptance tests, first you need to update the `script/gotestacc_vars.sh` with above variables values and then simply run `make testacc`.
```sh
$ make testacc
```




