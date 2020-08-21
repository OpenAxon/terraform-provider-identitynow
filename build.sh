#!/bin/bash
set -euo pipefail

go build -o terraform-provider-identitynow
mkdir -p ~/.terraform.d/plugins/registry.terraform.fake/axon/identitynow/0.1.0/darwin_amd64
mv terraform-provider-identitynow ~/.terraform.d/plugins/registry.terraform.fake/axon/identitynow/0.1.0/darwin_amd64/terraform-provider-identitynow_v0.1.0