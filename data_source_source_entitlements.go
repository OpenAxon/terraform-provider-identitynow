package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSourceEntitlement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSourceEntitlementRead,

		Schema: sourceEntitlementFields(),
	}
}

func dataSourceSourceEntitlementRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Getting Data source for Entitlements. Source ID %s", d.Get("source_id").(string))
	client, err := meta.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	sourceEntitlements, err := client.GetSourceEntitlement(context.Background(), d.Get("source_id").(string), d.Get("name").(string))
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Source ID %s not found.", d.Get("source_id").(string))
			return nil
		}
		return err
	}

	return flattenSourceEntitlement(d, sourceEntitlements[0])
}
