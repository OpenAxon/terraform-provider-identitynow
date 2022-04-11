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
	aggregationCompleted := false
	aggregationStart := func() error {

		//Ensure completed flag set regardless of outcome
		defer func() {
			aggregationCompleted = true
		}()

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
				return nil
			}
		}
	}

	var entitlement *Items
	for {
		entitlements, err := client.GetSourceEntitlements(context.Background(), d.Get("source_id").(string))
		if err != nil {
			return err
		}
		entitlement, err = getEntitlement(entitlements.Items, d.Get("name").(string))
		if err != nil {

			// When not found...
			if _, notFound := err.(NotFoundError); notFound {

				// ... agreggation will begin only if an aggregation source ID was provided and that the aggregation has not already
				// comptleted
				if aggregationSourceID != "" && !aggregationCompleted {
					log.Printf("[INFO] starting soure entitlement aggregation")
					err = aggregationStart()
					if err != nil {
						return err
					}
					continue
				}

				// The resource will be removed from state only if the aggregation source ID was not provided or it was but the
				// aggregation phase has completed.
				d.SetId("")
			}

			// Error finally returned if no data source was found
			return err
		}

		// This will be hit if the data source was found
		return flattenSourceEntitlement(d, entitlement)
	}
}
