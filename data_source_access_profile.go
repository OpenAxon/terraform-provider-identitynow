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

			"source": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: accessProfileSourceFields(),
				},
			},

			"owner": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceOwnerFields(),
				},
			},

			"entitlements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: accessProfileEntitlementsFields(),
				},
			},

			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"requestable": {
				Type:     schema.TypeBool,
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
