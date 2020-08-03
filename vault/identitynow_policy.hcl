# Necessary for Terraform, which creates short-lived tokens from the token you give it
path "auth/token/create" {
  capabilities = ["create", "read", "update", "list"]
}

# List and read key/value secrets
path "identitynow/*"
{
  capabilities = ["read", "list"]
}