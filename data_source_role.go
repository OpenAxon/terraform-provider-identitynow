package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_profile_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"approval_schemes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"denied_comments_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_comments_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"requestable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"revoke_request_approval_schemes": {
				Type:     schema.TypeString,
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
		_, notFound := err.(NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Role ID %s not found.", d.Get("id").(string))
			return nil
		}
		return err
	}

	return flattenRole(d, role)
}
