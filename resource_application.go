package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationUpdate,
		Delete: resourceApplicationDelete,

		Schema: applicationFields(),
	}
}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {
	application, err := expandApplication(d)
	if err != nil {
		return err
	}

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newApplication, err := client.CreateApplication(context.Background(), application)
	if err != nil {
		return err
	}

	err = flattenApplication(d, newApplication)
	if err != nil {
		return err
	}

	return resourceApplicationRead(d, m)
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	app := &Application{}
	err = client.GetApplication(context.Background(), d.Id(), app)
	if err != nil {
		return err
	}

	return flattenApplication(d, app)
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	app, err := expandApplication(d)
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}
	app.Alias = app.Name
	err = client.UpdateApplication(context.Background(), app.ID, app)
	if err != nil {
		return err
	}
	err = flattenApplication(d, app)
	if err != nil {
		return err
	}

	return resourceApplicationRead(d, m)
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	application, err := expandApplication(d)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	err = client.DeleteApplication(context.Background(), application.AppID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
