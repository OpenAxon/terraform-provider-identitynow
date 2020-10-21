package main

import (
	"context"
	"fmt"
	schema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceScheduleAccountAggregation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountAggregationScheduleCreateUpdate,
		Read:   resourceAccountAggregationScheduleRead,
		Update: resourceAccountAggregationScheduleCreateUpdate,
		Delete: resourceAccountAggregationScheduleDelete,

		Schema: accountAggregationScheduleFields(),
	}
}

func resourceAccountAggregationScheduleCreateUpdate(d *schema.ResourceData, m interface{}) error {
	accountAggregationSchedule, err := expandAccountAggregationSchedule(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Performing Account Aggregation Schedule for source ID %s", accountAggregationSchedule.SourceID)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newAccountAggregationSchedule, err := client.ManageAccountAggregationSchedule(context.Background(), accountAggregationSchedule, true)
	if err != nil {
		return err
	}

	newAccountAggregationSchedule.SourceID = accountAggregationSchedule.SourceID

	err = flattenAccountAggregationSchedule(d, newAccountAggregationSchedule)
	if err != nil {
		return err
	}

	return resourceAccountAggregationScheduleRead(d, m)
}

func resourceAccountAggregationScheduleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Refreshing Account Aggregation Schedule for source ID %s", d.Id())
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accountAggregationSchedule, err := client.GetAccountAggregationSchedule(context.Background(), d.Id())
	if accountAggregationSchedule.CronExpressions != nil {
		accountAggregationSchedule.SourceID = d.Id()
	}
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Account Aggregation Schedule for Source ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = flattenAccountAggregationSchedule(d, accountAggregationSchedule)
	if err != nil {
		return err
	}

	return nil
}

func resourceAccountAggregationScheduleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Deleting Account Aggregation for Source ID %s", d.Id())

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accountAggregationSchedule, err := client.GetAccountAggregationSchedule(context.Background(), d.Id())
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Account Aggregation Schedule for source ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	if accountAggregationSchedule.CronExpressions != nil {
		accountAggregationSchedule.SourceID = d.Id()
		_, err = client.ManageAccountAggregationSchedule(context.Background(), accountAggregationSchedule, false)
		if err != nil {
			return fmt.Errorf("Error removing Account Aggregation Schedule for source ID: %s. \nError: %s", d.Id(), err)
		}

		d.SetId("")
	}

	return nil
}
