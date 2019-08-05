package main

import (
	"log"

	"github.com/aareet/go-nfs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNFSDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceNFSDNSRecordCreate,
		Read:   resourceNFSDNSRecordRead,
		Delete: resourceNFSDNSRecordDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNFSDNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Organization).client
	domain := d.Get("domain").(string)
	params := map[string]string{
		"name": d.Get("name").(string),
		"type": d.Get("type").(string),
		"data": d.Get("data").(string),
	}
	resp, err := nfs.SetDNSRecord(client, domain, params)
	if err != nil {
		return err
	}
	d.SetId(params["name"])
	log.Printf("[DEBUG] ID: %s  and response: %s", d.Id(), resp)
	return resourceNFSDNSRecordRead(d, m)
}

func resourceNFSDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNFSDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Organization).client
	domain := d.Get("domain").(string)
	params := map[string]string{
		"name": d.Get("name").(string),
		"type": d.Get("type").(string),
		"data": d.Get("data").(string),
	}
	resp, err := nfs.RemoveDNSRecord(client, domain, params)
	if err != nil {
		return err
	}
	d.SetId(params["name"])
	log.Printf("[DEBUG] ID: %s  and response: %s", d.Id(), resp)
	return resourceNFSDNSRecordRead(d, m)
}
