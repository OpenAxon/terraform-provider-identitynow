package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,

		Schema: roleFields(),
	}
}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	role, err := expandRole(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating Role %s", role.Name)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newRole, err := client.CreateRole(context.Background(), role)
	if err != nil {
		return err
	}

	err = flattenRole(d, newRole)
	if err != nil {
		return err
	}

	return resourceRoleRead(d, m)
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Refreshing Role ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	role, err := client.GetRole(context.Background(), d.Id())
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(NotFoundError)
		if notFound {
			log.Printf("[INFO] Role ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = flattenRole(d, role)
	if err != nil {
		return err
	}

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Updating Role ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	log.Printf("disabled in role: %s\n", d.Get("disabled"))

	updatedRole, err := expandRole(d)
	log.Printf("role after expand: %v\n", updatedRole)
	if err != nil {
		return err
	}

	_, err = client.UpdateRole(context.Background(), updatedRole)
	if err != nil {
		return err
	}

	return resourceRoleRead(d, m)
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Deleting Role ID %s", d.Id())

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	role, err := client.GetRole(context.Background(), d.Id())
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(NotFoundError)
		if notFound {
			log.Printf("[INFO] Role ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	_, err = client.DeleteRole(context.Background(), role)
	if err != nil {
		return fmt.Errorf("error removing Role: %s", err)
	}

	d.SetId("")
	return nil
}
