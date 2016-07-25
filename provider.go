package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jsok/terraform-provider-dynadot/dynadot"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.dynadot.com/api3.xml",
				Description: "Dynadot v3 API URL",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNADOT_API_KEY", nil),
				Description: "Dynadot API key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dynadot_domain_nameservers": resourceDynadotDomainNameservers(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return dynadot.NewClient(d.Get("api_url").(string), d.Get("api_key").(string))
}
