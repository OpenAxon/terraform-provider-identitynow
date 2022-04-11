package main

import (
	"context"
	"log"
	"time"

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

	aggregationSourceID := d.Get("aggregation_source_id").(string)
	sourceEntitlements, err := client.GetSourceEntitlements(context.Background(), d.Get("source_id").(string))
	if err != nil {
		return err
	}
	entitlement, err := getEntitlement(sourceEntitlements.Items, d.Get("name").(string))
	if _, notFound := err.(NotFoundError); notFound {
		if aggregationSourceID != "" {

			log.Printf("[INFO] starting soure entitlement aggregation")

			result, err := client.StartSourceEntitlementAggregation(context.Background(), aggregationSourceID)
			if err != nil {
				return err
			}

			for {
				log.Printf("[INFO] polling aggregation task %s", result.Task.ID)
				time.Sleep(2 * time.Second)

				status, err := client.GetSourceEntitlementAggregationStatus(context.Background(), result.Task.ID)
				if err != nil {
					return err
				}

				log.Printf("[INFO] task completion status '%s'", status.CompletionStatus)

				if status.CompletionStatus == "Success" {
					// This sleep is required for the API objects to become available post aggregation.
					time.Sleep(2 * time.Second)
					break
				}
			}

			sourceEntitlements, err = client.GetSourceEntitlements(context.Background(), d.Get("source_id").(string))
			if err != nil {
				return err
			}

			entitlement, err = getEntitlement(sourceEntitlements.Items, d.Get("name").(string))
			if err != nil {
				return err
			}

		} else {
			return err
		}
	} else {
		return err
	}

	return flattenSourceEntitlement(d, entitlement)
}
