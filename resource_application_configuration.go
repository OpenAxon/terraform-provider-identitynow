package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApplicationConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationConfigurationCreate,
		Read:   resourceApplicationConfigurationRead,
		Update: resourceApplicationConfigurationUpdate,
		Delete: resourceApplicationConfigurationDelete,

		Schema: applicationConfigurationFields(),
	}
}

func resourceApplicationConfigurationCreate(d *schema.ResourceData, m interface{}) error {

	config, err := expandApplicationConfiguration(d)
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	defaultConfig := &ApplicationConfiguration{}
	err = client.GetApplication(context.Background(), config.ID, defaultConfig)
	if err != nil {
		return err
	}

	// Intentional - this resource controls a subset of application properties
	err = client.UpdateApplication(context.Background(), config.ID, config)
	if err != nil {
		return err
	}

	err = flattenApplicationConfiguration(d, config)
	if err != nil {
		return err
	}

	// This is done for creation to store the original app configuration prior to provisioning this resource
	// which is what it will be retured to when this resource is deleted

	defaultValues, _ := json.Marshal(defaultConfig)
	d.Set("default_values", string(defaultValues))

	time.Sleep(time.Millisecond * 500)
	return resourceApplicationConfigurationRead(d, m)
}

func resourceApplicationConfigurationRead(d *schema.ResourceData, m interface{}) error {
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	config := &ApplicationConfiguration{}
	err = client.GetApplication(context.Background(), d.Id(), config)
	if err != nil {
		return err
	}

	return flattenApplicationConfiguration(d, config)
}

func resourceApplicationConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	config, err := expandApplicationConfiguration(d)
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	// Intentional - this resource controls a subset of application properties
	err = client.UpdateApplication(context.Background(), config.ID, config)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)
	return resourceApplicationConfigurationRead(d, m)
}

func resourceApplicationConfigurationDelete(d *schema.ResourceData, m interface{}) error {

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	defaultValues := d.Get("default_values").(string)
	config := &ApplicationConfiguration{}
	json.Unmarshal([]byte(defaultValues), config)

	err = client.UpdateApplication(context.Background(), config.ID, config)
	if err != nil {
		return err
	}
	d.SetId("")
	time.Sleep(time.Millisecond * 500)
	return nil
}
