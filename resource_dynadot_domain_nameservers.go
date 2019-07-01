package main

import (
	"strings"

	"github.com/codewise/terraform-provider-dynadot/dynadot"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDynadotDomainNameservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceDynadotDomainNameserversCreate,
		Read:   resourceDynadotDomainNameserversRead,
		Update: resourceDynadotDomainNameserversUpdate,
		Delete: resourceDynadotDomainNameserversDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nameservers": &schema.Schema{
				Type:     schema.TypeList,
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

	aInterface := d.Get("nameservers").([]interface{})
	nameservers := make([]string, len(aInterface))
	for i, ns := range aInterface {
		nameservers[i] = ns.(string)
	}

	err := client.SetDomainNameServers(domain, nameservers)
	if err != nil {
		return err
	}

	d.SetId(toId(nameservers, domain))
	return resourceDynadotDomainNameserversRead(d, meta)
}

func toId(nameservers []string, domain string) string {
	return strings.Join(append(nameservers, domain), "_")
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

func resourceDynadotDomainNameserversDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(dynadot.Client)
	domain := d.Get("domain").(string)

	err := client.ParkDomain(domain)
	if err != nil {
		return err
	}

	return nil
}
