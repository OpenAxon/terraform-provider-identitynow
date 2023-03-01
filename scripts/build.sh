#!/bin/bash
set -euo pipefail

# Set the version to something with the format x.y.z, ideally not a version already published externally
# Once the directory ~/.terraform.d/plugins/registry.terraform.io/openaxon/identitynow is created,
# TF will not longer attempt to look for any versions of this provider on the remote registry.
# When finished with development, you can remove the folder from your laptop to start using the public provider again.
# https://www.terraform.io/docs/commands/cli-config.html#implied-local-mirror-directories
VERSION=0.3.1
go build -o terraform-provider-identitynow
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/openaxon/identitynow/${VERSION}/darwin_amd64
mv terraform-provider-identitynow ~/.terraform.d/plugins/registry.terraform.io/openaxon/identitynow/${VERSION}/darwin_amd64/terraform-provider-identitynow_v${VERSION}