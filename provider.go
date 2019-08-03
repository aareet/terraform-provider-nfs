package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//Provider NFSN provider to manage DNS records
func Provider() terraform.ResourceProvider {
	var p *schema.Provider
	p = &schema.Provider{

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nfs_dns_record": resourceNFSDNSRecord(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)

	return p
}
func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			accountID: d.Get("account_id").(string),
			apiKey:    d.Get("api_key").(string),
			login:     d.Get("login").(string),
		}

		meta, err := config.Client()
		if err != nil {
			return nil, err
		}

		return meta, nil
	}
}
