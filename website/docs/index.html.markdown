---
layout: "identitynow"
page_title: "Provider: IdentityNow"
description: |-
  The IdentityNow Provider is used to interact with the many resources supported by [IdentityNow API](https://developer.sailpoint.com/idn/api/v3).

---

# IdentityNow Provider

The IdentityNow Provider can be used to configure infrastructure in [SailPoint IdentitNow](https://www.sailpoint.com/products/identitynow/) using the official API's. Documentation regarding the [Data Sources](/docs/configuration/data-sources.html) and [Resources](/docs/configuration/resources.html) supported by the IdentityNow Provider can be found in the navigation to the left.

Interested in the provider's latest features, or want to make sure you're up to date? Check out the [changelog](https://github.com/OpenAxon/terraform-provider-identitynow/blob/master/CHANGELOG.md) for version information and release notes.

## Authenticating to IdentityNow

The IdentityNow Provider follows the [Client Credentials Grant Flow](https://developer.sailpoint.com/idn/api/authentication/#client-credentials-grant-flow), using the Client ID and Client Secret obtained from the personal access token.

## Example Usage

```hcl
# We strongly recommend using the required_providers block to set the
# IdentityNow Provider source and version being used
terraform {
  required_providers {
    identitynow = {
      source  = "OpenAxon/identitynow"
      version = "=0.3.1"
    }
  }
}

# Configure the IdentityNow Provider
provider "identitynow" {
  api_url       = "<org_name>.api.identitynow.com"
  client_id     = "<client_id>"
  client_secret = "<clien_secret>"
}

# Create a source
resource "identitynow_source" "example" {
  name     = "example-resources"
  # ...
}
```

## Features and Bug Requests

The IdentityNow provider's bugs and feature requests can be found in the [GitHub repo issues](https://github.com/OpenAxon/terraform-provider-identitynow/issues).
Please avoid "me too" or "+1" comments. Instead, use a thumbs up [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/)
on enhancement requests. Provider maintainers will often prioritize work based on the number of thumbs on an issue.

Community input is appreciated on outstanding issues! We love to hear what use
cases you have for new features, and want to provide the best possible
experience for you using the IdentityNow provider.

If you have a bug or feature request without an existing issue

* if an existing resource or field is working in an unexpected way, [file a bug](https://github.com/OpenAxon/terraform-provider-identitynow/issues/new?template=Bug_Report.md).

* if you'd like the provider to support a new resource or field, [file an enhancement/feature request](https://github.com/OpenAxon/terraform-provider-identitynow/issues/new?template=Feature_Request.md).

The provider maintainers will often use the assignee field on an issue to mark
who is working on it.

* An issue assigned to an individual maintainer indicates that the maintainer is working
on the issue

* If you're interested in working on an issue please leave a comment on that issue


## Argument Reference

The following arguments are supported:

* `api_url` - (Required) The URL to the IdentityNow API.

* `client_id` - (Required) API client used to authenticate with the IdentityNow API.

* `client_secret` - (Required) API client secret used to authenticate with the IdentityNow API.The Cloud Environment which should be used. Possible values are `public`, `usgovernment`, `german`, and `china`. Defaults to `public`. This can also be sourced from the `ARM_ENVIRONMENT` Environment Variable.
