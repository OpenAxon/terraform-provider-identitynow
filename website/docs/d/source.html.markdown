---
subcategory: "Source"
layout: "identitynow"
page_title: "IdentityNow: Data Source: identitynow_source"
description: |-
  Gets information about an existing Source.
---

# Data Source: identitynow_source

Use this data source to access information about an existing Source.

## Example Usage

```hcl
data "identitynow_source" "example" {
  name = "example"
}

output "identitynow_source_description" {
  value = data.identitynow_source.description
}
```

## Arguments Reference

The following arguments are supported:

* TBU

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* TBU

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the Source.