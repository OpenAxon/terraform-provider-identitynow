package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"accessProfiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: roleAccessProfilesFields(),
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceOwnerFields(),
				},
			},
			"requestable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Getting Data source for Role ID %s", d.Get("id").(string))
	client, err := meta.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	role, err := client.GetRole(context.Background(), d.Get("id").(string))
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Role ID %s not found.", d.Get("id").(string))
			return nil
		}
		return err
	}

	return flattenRole(d, role)
}
