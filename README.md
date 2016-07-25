[![Build Status](https://travis-ci.org/jsok/terraform-provider-dynadot.svg?branch=master)](https://travis-ci.org/jsok/terraform-provider-dynadot)

# Terraform Dynadot provider

A Terraform provider plugin to manage nameservers for domains registered
with [Dynadot](https://www.dynadot.com/).

Makes use of the [Advanced Domain API](https://www.dynadot.com/domain/api3.html).

## Provider

### Example Usage

```hcl
# Configure the Dynadot provider
provider "dynadot" {
    api_key = "secret"
}
```

### Argument Reference

The following arguments are supported:

 * `api_key` - (*Required*) Your API key. It must be provided, but it can also
   be sourced from the `DYNADOT_API_KEY` environment variable.
   See https://www.dynadot.com/community/help/question/find-API-settings for
   details on how to retrieve your key.

 * `api_url` - The base URL of the Dynadot API.
   Default: https://api.dynadot.com/api3.xml

## Resources

### `dynadot_domain_nameservers`

Set the nameservers of a domain.

```hcl
resource "dynadot_domain_nameservers" "example_com" {
    domain = "example.com"
    nameservers = [
        "ns1.example.com",
        "ns2.example.com"
    ]
}
```

### Argument Reference

The following arguments are supported:

 * `domain` - (*Required*) The domain name you want to set nameservers for.
 * `nameservers` - (*Required*) A string list of nameservers, maximum of 13
   nameservers.

The domain you are attempting to set nameservers for must have the
"Name Server Setting" set to "Name Servers", if it is set to something else
e.g. "Dynadot Parking", this resource will fail.
