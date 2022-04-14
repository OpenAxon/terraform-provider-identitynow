package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGovernanceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGovernanceGroupCreate,
		Read:   resourceGovernanceGroupRead,
		Update: resourceGovernanceGroupUpdate,
		Delete: resourceGovernanceGroupDelete,
		Schema: governanceGroupFields(),
	}
}

func resourceGovernanceGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("resourceGovernanceGroupCreate")
	governanceGroup, err := expandGovernanceGroup(d)
	if err != nil {
		return err
	}

	c, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newGovernanceGroup, err := c.CreateGovernanceGroup(context.Background(), *governanceGroup)
	if err != nil {
		return err
	}

	err = flattenGovernanceGroup(d, newGovernanceGroup)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)
	return resourceGovernanceGroupRead(d, m)

}

func resourceGovernanceGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("resourceGovernanceGroupRead")
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	governanceGroup, err := client.GetGovernanceGroup(context.Background(), d.Id())
	if err != nil {
		if _, notFound := err.(NotFoundError); notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	err = flattenGovernanceGroup(d, governanceGroup)
	if err != nil {
		return err
	}

	return nil
}

func resourceGovernanceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("resourceGovernanceGroupUpdate")
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	governanceGroup, err := expandGovernanceGroup(d)
	if err != nil {
		return err
	}

	_, err = client.UpdateGovernanceGroup(context.Background(), *governanceGroup)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)
	return resourceGovernanceGroupRead(d, m)
}

func resourceGovernanceGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("resourceGovernanceGroupDelete")
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	governanceGroup, err := client.GetGovernanceGroup(context.Background(), d.Id())
	if err != nil {
		if _, notFound := err.(NotFoundError); notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	err = client.DeleteGovernanceGroup(context.Background(), governanceGroup.ID)
	if err != nil {
		return err
	}

	d.SetId("")
	time.Sleep(time.Millisecond * 500)
	return nil
}
