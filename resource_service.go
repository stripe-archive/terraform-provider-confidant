package main

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stripe/go-confidant-client/confidant"
)

func getCredentialsList(d *schema.ResourceData) ([]string, error) {
	set, ok := d.Get("credentials").(*schema.Set)
	if !ok {
		return nil, errors.New("Could not cast credentials to *schema.Set")
	}
	interfaces := set.List()
	strings := make([]string, 0, len(interfaces))
	for _, item := range interfaces {
		credentialName, ok := item.(string)
		if !ok {
			return nil, errors.New("Could not cast credential name to string")
		}
		strings = append(strings, credentialName)
	}
	return strings, nil
}

func resourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,
		Exists: resourceServiceExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"credentials": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

// resourceServiceCreate calls CreateService with no credentials.
// If the service already exists, it sets enabled to true and updates the credentials.
func resourceServiceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(confidant.Client)
	name, ok := d.Get("name").(string)
	if !ok {
		return errors.New("Could not cast service name to string")
	}
	credentials, err := getCredentialsList(d)
	if err != nil {
		return err
	}
	_, err = client.CreateService(name, credentials)
	if err != nil {
		if err.Error() != "Service Already Exists" {
			return err
		}
		if _, err = client.SetServiceCredentials(name, credentials); err != nil {
			return err
		}
		_, err := client.EnableService(name)
		if err != nil {
			return err
		}
	}
	d.SetId(name)
	return nil
}

func resourceServiceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(confidant.Client)
	name := d.Id()

	service, err := client.GetService(name)
	if err != nil {
		return err
	}
	credentials := make([]string, 0, len(service.Credentials))
	for _, credential := range service.Credentials {
		credentials = append(credentials, credential.Name)
	}
	if !service.Enabled {
		return nil
	}
	d.Set("name", name)
	d.Set("credentials", credentials)
	return nil
}

// resourceServiceUpdate sets credentials for the service
func resourceServiceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(confidant.Client)
	name, ok := d.Get("name").(string)
	if !ok {
		return errors.New("Could not cast service name to string")
	}
	credentials, err := getCredentialsList(d)
	if err != nil {
		return err
	}
	if _, err = client.SetServiceCredentials(name, credentials); err != nil {
		return err
	}
	if _, err := client.EnableService(name); err != nil {
		return err
	}
	return nil
}

// resourceServiceDelete updates the service with enabled set to false
func resourceServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(confidant.Client)
	name, ok := d.Get("name").(string)
	if !ok {
		return errors.New("Could not cast service name to string")
	}
	if _, err := client.DisableService(name); err != nil {
		return err
	}
	d.Set("credentials", nil)
	d.Set("enabled", false)
	return nil
}

// resourceServiceExists checks if the service is enabled
func resourceServiceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(confidant.Client)
	name := d.Id()

	service, err := client.GetService(name)
	if err != nil {
		if err.Error() == "Service Doesn't Exist" {
			return false, nil
		}
		// Some other error, default to true
		return true, err
	}
	return service.Enabled, nil
}
