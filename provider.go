package main

import (
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stripe/go-confidant-client/confidant"
	"github.com/stripe/go-confidant-client/kmsauth"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"authkey": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confidant KMS key",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confidant URL",
			},
			"to": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confidant IAM role",
			},
			"from": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS user id",
			},
			"unixproxy": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional unix socket proxy",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS region (for KMS encrypt request)",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"confidant_service": resourceService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	authkey := d.Get("authkey").(string)
	to := d.Get("to").(string)
	from := d.Get("from").(string)
	region := d.Get("region").(string)
	userType := "user"
	url := d.Get("url").(string)
	httpClient := &http.Client{}
	proxy, ok := d.GetOk("unixproxy")
	if ok {
		httpClient.Transport = confidant.UnixProxy(proxy.(string))
	}
	generator := kmsauth.NewTokenGenerator(authkey, to, from, userType, region)
	client := confidant.NewClient(url, httpClient, &generator)
	return client, nil
}
