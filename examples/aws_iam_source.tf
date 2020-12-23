resource "identitynow_source" "aws_iam_source" {
  name        = "AWS IAM Source"
  description = "The AWS IAM connector created by terraform"
  connector   = "aws"

  owner {
    id   = data.identitynow_identity.john_doe.external_id
    name = data.identitynow_identity.john_doe.name
    type = "IDENTITY"
  }

  cluster {
    id   = "<CLUSTER_ID>"
    name = "<CLUSTER_NAME>"
  }

  connector_attributes {
    include_aws_account_id_list = "123456789012" #list of aws accounts id separated with `,`
    kid                         = "<ACCESS_KEY_ID>"
    secret                      = "ACCESS_KEY_SECRET>"
    role_name                   = "<AWS_ROLE_NAME>"
    connector_class             = "openconnector.connector.aws.AWSConnectorSDK"
    encrypted                   = "secret"
  }

  lifecycle {
    ignore_changes = [
      connector_attributes[0].secret,
      owner[0].name,
      management_workgroup[0].name,
    ]
  }
}
