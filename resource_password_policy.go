package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourcePasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePasswordPolicyCreate,
		Read:   resourcePasswordPolicyRead,
		Update: resourcePasswordPolicyUpdate,
		Delete: resourcePasswordPolicyDelete,

		Schema: passwordPolicyFields(),
	}

}

func resourcePasswordPolicyCreate(d *schema.ResourceData, m interface{}) error {
	passwordPolicy, err := expandPasswordPolicy(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating Password Policy %s", passwordPolicy.Name)

	c, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newPasswordPolicy, err := c.CreatePasswordPolicy(context.Background(), passwordPolicy)
	if err != nil {
		return err
	}

	err = flattenPasswordPolicy(d, newPasswordPolicy)
	if err != nil {
		return err
	}

	return resourcePasswordPolicyRead(d, m)

}

func resourcePasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Refreshing Password Policy ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	passwordPolicy, err := client.GetPasswordPolicy(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Password Policy ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = flattenPasswordPolicy(d, passwordPolicy)
	if err != nil {
		return err
	}

	return nil
}

func resourcePasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Updating Password Policy ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	updatedPasswordPolicy, err := expandPasswordPolicy(d)
	if err != nil {
		return err
	}

	_, err = client.UpdatePasswordPolicy(context.Background(), updatedPasswordPolicy)
	if err != nil {
		return err
	}

	return resourcePasswordPolicyRead(d, m)
}

func resourcePasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Deleting Password Policy ID %s", d.Id())

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	passwordPolicy, err := client.GetPasswordPolicy(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Password Policy ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = client.DeletePasswordPolicy(context.Background(), passwordPolicy.ID)
	if err != nil {
		return fmt.Errorf("error removing Passwprd Policy: %s", err)
	}

	d.SetId("")
	return nil
}
