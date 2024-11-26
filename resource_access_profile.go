package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceAccessProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessProfileCreate,
		Read:   resourceAccessProfileRead,
		Update: resourceAccessProfileUpdate,
		Delete: resourceAccessProfileDelete,

		Schema: accessProfileFields(),
	}
}

func resourceAccessProfileCreate(d *schema.ResourceData, m interface{}) error {
	accessProfile, err := expandAccessProfile(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating Access Profile %s", accessProfile.Name)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newAccessProfile, err := client.CreateAccessProfile(context.Background(), accessProfile)
	if err != nil {
		return err
	}

	err = flattenAccessProfile(d, newAccessProfile)
	if err != nil {
		return err
	}

	return resourceAccessProfileRead(d, m)
}

func resourceAccessProfileRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Refreshing Access Profile ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accessProfile, err := client.GetAccessProfile(context.Background(), d.Id())
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Access Profile ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = flattenAccessProfile(d, accessProfile)
	if err != nil {
		return err
	}

	return nil
}

func resourceAccessProfileUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Updating Access Profile ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	updatedAccessProfile, id, err := expandUpdateAccessProfile(d)
	if err != nil {
		return err
	}

	_, err = client.UpdateAccessProfile(context.Background(), updatedAccessProfile, id)
	if err != nil {
		return err
	}

	return resourceAccessProfileRead(d, m)
}

func resourceAccessProfileDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Deleting Access Profile ID %s", d.Id())

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accessProfile, err := client.GetAccessProfile(context.Background(), d.Id())
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Access Profile ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = client.DeleteAccessProfile(context.Background(), accessProfile)
	if err != nil {
		return fmt.Errorf("Error removing Access Profile: %s", err)
	}

	d.SetId("")
	return nil
}
