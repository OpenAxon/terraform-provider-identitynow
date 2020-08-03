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





