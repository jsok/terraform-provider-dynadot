package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jsok/terraform-provider-dynadot/dynadot"
)

func resourceDynadotDomainNameservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceDynadotDomainNameserversCreate,
		Read:   resourceDynadotDomainNameserversRead,
		Update: resourceDynadotDomainNameserversUpdate,
		Delete: resourceDynadotDomainNameserversUpdate,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nameservers": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MaxItems: 13,
			},
		},
	}
}

func resourceDynadotDomainNameserversCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(dynadot.Client)
	domain := d.Get("domain").(string)
	nameservers := d.Get("nameservers").([]string)

	// First determine if the nameservers are registered in the global list
	registered, err := client.ListNameServers()
	if err != nil {
		return nil
	}

	// Add any missing nameservers to the global list
	missing := make([]string, 0, len(nameservers))
	for _, ns := range nameservers {
		found := false
		for _, reg := range registered {
			if ns == reg {
				found = true
			}
		}
		if found == false {
			missing = append(missing, ns)
		}
	}

	for _, ns := range missing {
		if err := client.AddNameServer(ns); err != nil {
			return err
		}
	}

	d.SetId(strings.Join(append(nameservers, domain), "_"))

	// Set the domain's nameserver list
	return client.SetDomainNameServers(domain, nameservers)
}

func resourceDynadotDomainNameserversRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(dynadot.Client)

	domain := d.Get("domain").(string)

	info, err := client.GetDomainInfo(domain)
	if err != nil {
		return err
	}

	d.Set("nameservers", info.NameServers)

	return nil
}

func resourceDynadotDomainNameserversUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDynadotDomainNameserversCreate(d, meta)
}
