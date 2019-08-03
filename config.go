package main

import (
	"github.com/aareet/go-nfs"
)

//Config contains configuration data for the NFSN API client
type Config struct {
	accountID string
	apiKey    string
	login     string
}

//Organization contains the API client and can be configured with an account name
type Organization struct {
	name   string
	client *nfs.Client
}

//Client returns an initialized NFSN API client
func (c *Config) Client() (interface{}, error) {
	var org Organization
	org.name = "NFS Client for Terraform"
	org.client = nfs.NewClientForAccount(c.accountID, c.apiKey, c.login)
	return &org, nil
}
