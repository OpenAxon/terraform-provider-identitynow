package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccessProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccessProfileRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source id",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access Profile name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access Profile description",
			},

			"source_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Source Id that Access Profile is going to create for",
			},

			"source_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"owner_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"entitlements": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Access Profile Entitlements.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"denied_comments_required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Access Profile Denied Comments Required",
			},

			"approval_schemes": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"protected": {
				Type:     schema.TypeBool,
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

func dataSourceAccessProfileRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Data source for Access Profile ID %s", d.Get("id").(string))
	client, err := meta.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accessProfile, err := client.GetAccessProfile(context.Background(), d.Get("id").(string))
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Access Profile ID %s not found.", d.Get("id").(string))
			return nil
		}
		return err
	}

	return flattenAccessProfile(d, accessProfile)
}
