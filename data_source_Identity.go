package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceIdentity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIdentityRead,

		Schema: identityFields(),
	}
}

func dataSourceIdentityRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Getting Data source for Identity. Identity alias %s", d.Get("alias").(string))
	client, err := meta.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	identity, err := client.GetIdentity(context.Background(), d.Get("alias").(string))
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("[INFO] Data source for Identity alias %s not found.", d.Get("alias").(string))
			return nil
		}
		return err
	}

	return flattenIdentity(d, identity[0])
}
