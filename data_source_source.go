package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSourceRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source id",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source description",
			},
			"delete_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"authoritative": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if this source is authoritative",
			},
			"owner": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceOwnerFields(),
				},
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceSchemaFields(),
				},
			},
			"cluster": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceClusterFields(),
				},
			},
			"account_correlation_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceAccountCorrelationConfigFields(),
				},
			},
			"connector_attributes": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sourceConnectorAttributesFields(),
				},
			},
		},
	}
}

func dataSourceSourceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Data source for Source ID %s", d.Get("id").(string))
	client, err := meta.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	source, err := client.GetSource(context.Background(), d.Get("id").(string))
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Source ID %s not found.", d.Get("id").(string))
			return nil
		}
		return err
	}

	return flattenSourceAAD(d, source)
}
