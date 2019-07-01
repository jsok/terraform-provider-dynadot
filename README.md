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

Set the nameservers of a domain. Upon destroying this resource the domain's DNS settings will be set to "Dynadot Parking".

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
   
   
## Plugin installation guide

[As of now](https://github.com/hashicorp/terraform/issues/15252) Terraform doesn't support third party plugin distribution servers.
 
Currently to install a custom or unsupported third-party binary, you have a few options:

* Install the provider binary to $HOME/.terraform.d/plugins (or windows equivalent) ([documented here](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins)).
    * if you're using Linux, `make install` will do exactly that
* Preinstall the binary on the local filesystem, and use the -plugin-dir flag during initialisation ([documented here](https://learn.hashicorp.com/terraform/development/running-terraform-in-automation#terraform-init-input-false-plugin-dir-usr-lib-custom-terraform-plugins)).
* Commit the .terraform/plugins directory within the project. The resultant directory tree should look like that: 
    ```text
    .
    ├── main.tf
    ├── .terraform
    │   └── plugins
    │       └── linux_amd64
    │           ├── lock.json
    │           └── terraform-provider-dynadot
    ├── terraform.tfstate
    └── terraform.tfstate.backup
    
    ```
