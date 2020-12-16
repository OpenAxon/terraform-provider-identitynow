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
Edit the Go files that make up the provider, and rebuild the provider.
```bash
./build.sh
```
This script places the provider binary in an implied local mirror directory ($HOME/.terraform.d/plugins/). See build.sh
for more comments about ensuring that Terraform uses the local mirror rather than searching the remote registry. 

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

Note: Due to a bug in IdentityNow, Encrypted field in ConnectorAttributes block cannot be left null.




