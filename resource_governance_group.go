package main

import (
	"context"
	"fmt"

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

	return resourceGovernanceGroupRead(d, m)

}

func resourceGovernanceGroupRead(d *schema.ResourceData, m interface{}) error {
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	governanceGroup, err := client.GetGovernanceGroup(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
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

	return resourceGovernanceGroupRead(d, m)
}

func resourceGovernanceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	governanceGroup, err := client.GetGovernanceGroup(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	res, err := client.DeleteGovernanceGroup(context.Background(), governanceGroup.ID)
	if err != nil {
		return err
	}

	if len(res.Deleted) != 1 {
		return fmt.Errorf("expected result id array to be 1, got %d :%v", len(res.Deleted), res.Deleted)
	}

	if res.Deleted[0] != governanceGroup.ID {
		return fmt.Errorf("expected result id to be %s, got %s", governanceGroup.ID, res.Deleted[0])
	}

	d.SetId("")
	return nil
}
