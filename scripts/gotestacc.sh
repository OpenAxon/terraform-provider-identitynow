#!/usr/bin/env bash

set -e

echo "==>Running acceptance testing..."

source $(dirname $0)/gotestacc_vars.sh

TF_ACC=1 go test -cover -tags=test ./... -v -timeout 120m