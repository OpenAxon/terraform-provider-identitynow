package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceAccountSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountSchemaCreate,
		Read:   resourceAccountSchemaRead,
		Update: resourceAccountSchemaUpdate,
		Delete: resourceAccountSchemaDelete,
		Importer: &schema.ResourceImporter{
			State: resourceAccountSchemaImport,
		},

		Schema: accountSchemaFields(),
	}
}

func resourceAccountSchemaCreate(d *schema.ResourceData, m interface{}) error {
	accountSchema, err := expandAccountSchema(d)
	if err != nil {
		return err
	}

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newAccountSchema, err := client.GetAccountSchema(context.Background(), accountSchema.SourceID, accountSchema.ID)
	if err != nil {
		// Handle NotFoundError and other errors as before
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("Source ID %s not found.", accountSchema.SourceID)
			d.SetId("")
			return nil
		}
		return err
	}

	newAccountSchema.SourceID = accountSchema.SourceID

	// Use a map to track seen attributes based on a combination of fields
	seen := make(map[string]bool)
	result := []*AccountSchemaAttribute{}

	// Iterate over accountSchema.Attributes to filter out duplicates
	for _, attribute := range accountSchema.Attributes {
		// Create a unique key for each attribute based on important fields (e.g., "name")
		key := attribute.Name // or use a combination of "Name" and other unique fields if necessary
		if _, ok := seen[key]; !ok {
			seen[key] = true
			result = append(result, attribute)
		}
	}

	// Update the newAccountSchema with the filtered attributes
	newAccountSchema.Attributes = result
	newAccountSchema.ID = accountSchema.ID

	log.Printf("[INFO] Creating Account Schema Attribute for source %+v\n", newAccountSchema.SourceID)

	accountSchemaResponse, err := client.UpdateAccountSchema(context.Background(), newAccountSchema)
	if err != nil {
		return err
	}

	accountSchemaResponse.SourceID = accountSchema.SourceID

	err = flattenAccountSchema(d, accountSchemaResponse)
	if err != nil {
		return err
	}

	return resourceAccountSchemaRead(d, m)
}

func resourceAccountSchemaRead(d *schema.ResourceData, m interface{}) error {
	sourceId := d.Get("source_id").(string)
	schemaId := d.Get("schema_id").(string)
	attrName := d.Get("name").(string)
	log.Printf("[INFO] Refreshing Account Schema for Source %s", sourceId)
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accountSchema, err := client.GetAccountSchema(context.Background(), sourceId, schemaId)
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("Source ID %s not found.", sourceId)
			d.SetId("")
			return nil
		}
		return err
	}
	if accountSchema.Attributes == nil {
		log.Printf("Attribute %s not found in Account Schema.", attrName)
		d.SetId("")
	}

	accountSchema.SourceID = sourceId
	err = flattenAccountSchema(d, accountSchema)
	if err != nil {
		return err
	}

	return nil
}

func resourceAccountSchemaUpdate(d *schema.ResourceData, m interface{}) error {
	updatedAccountSchema, err := expandAccountSchema(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Updating %s for Account Schema for source ID %s", d.Get("name").(string), d.Get("source_id").(string))
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	_, err = client.UpdateAccountSchema(context.Background(), updatedAccountSchema)
	if err != nil {
		return err
	}

	return resourceAccountSchemaRead(d, m)
}

func resourceAccountSchemaDelete(d *schema.ResourceData, m interface{}) error {
	sourceId := d.Get("source_id").(string)
	schemaId := d.Get("schema_id").(string)
	name := d.Get("name").(string)
	log.Printf("[INFO] Deleting %s from Account Schema for source ID %s", name, sourceId)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	accountSchema, err := client.GetAccountSchema(context.Background(), sourceId, schemaId)
	if err != nil {
		// non-panicking type assertion, 2nd arg is boolean indicating type match
		_, notFound := err.(*NotFoundError)
		if notFound {
			log.Printf("Source ID %s not found.", sourceId)
			d.SetId("")
			return nil
		}
		return err
	}

	if accountSchema.Attributes == nil {
		log.Printf("Attribute %s not found in Account Schema.", name)
		d.SetId("")
	}

	accountSchema.SourceID = sourceId

	err = client.DeleteAccountSchema(context.Background(), accountSchema)
	if err != nil {
		return fmt.Errorf("error removing Account Schema from source %s. Error: %s", sourceId, err)
	}

	d.SetId("")
	return nil
}
