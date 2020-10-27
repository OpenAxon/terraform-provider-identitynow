package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceAccountSchemaAttribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountSchemaAttributeCreate,
		Read:   resourceAccountSchemaAttributeRead,
		Update: resourceAccountSchemaAttributeUpdate,
		Delete: resourceAccountSchemaAttributeDelete,
		Importer: &schema.ResourceImporter{
			State: resourceAccountSchemaAttributeImport,
		},

		Schema: accountSchemaAttributeFields(),
	}
}

func resourceAccountSchemaAttributeCreate(d *schema.ResourceData, m interface{}) error {
	attribute, err := expandAccountSchemaAttribute(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating Account Schema Attribute %s", attribute.Name)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	newAttribute, err := client.CreateAccountSchemaAttribute(context.Background(), attribute)
	if err != nil {
		return err
	}

	newAttribute.ObjectType = attribute.ObjectType
	newAttribute.SourceID = attribute.SourceID

	err = flattenAccountSchemaAttribute(d, newAttribute)
	if err != nil {
		return err
	}

	return resourceAccountSchemaAttributeRead(d, m)
}

func resourceAccountSchemaAttributeRead(d *schema.ResourceData, m interface{}) error {
	sourceId := d.Get("source_id").(string)
	attrName := d.Get("name").(string)
	log.Printf("[INFO] Refreshing Account Schema for Source %s", sourceId)
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	attributes, err := client.GetAccountSchemaAttributes(context.Background(), sourceId)
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
	attribute := getAccountSchemaAttribute(attributes, attrName)
	if attribute == nil {
		log.Printf("Attribute %s not found in Account Schema.", attrName)
		d.SetId("")
	}

	attribute.SourceID = sourceId
	attribute.ObjectType = attributes.ObjectType

	err = flattenAccountSchemaAttribute(d, attribute)
	if err != nil {
		return err
	}

	return nil
}

func resourceAccountSchemaAttributeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO] Updating Attribute %s for Account Schema for source ID %s", d.Get("name").(string), d.Get("source_id").(string))
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	updatedAttribute, err := expandAccountSchemaAttribute(d)
	if err != nil {
		return err
	}

	_, err = client.UpdateAccountSchemaAttribute(context.Background(), updatedAttribute)
	if err != nil {
		return err
	}

	return resourceAccountSchemaAttributeRead(d, m)
}

func resourceAccountSchemaAttributeDelete(d *schema.ResourceData, m interface{}) error {
	sourceId := d.Get("source_id").(string)
	attrName := d.Get("name").(string)
	log.Printf("[INFO] Deleting Attribute %s from Account Schema for source ID %s", attrName, sourceId)

	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	attributes, err := client.GetAccountSchemaAttributes(context.Background(), sourceId)
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
	attribute := getAccountSchemaAttribute(attributes, attrName)

	if attribute == nil {
		log.Printf("Attribute %s not found in Account Schema.", attrName)
		d.SetId("")
	}

	attribute.ObjectType = attributes.ObjectType
	attribute.SourceID = sourceId

	_, err = client.DeleteAccountSchemaAttribute(context.Background(), attribute)
	if err != nil {
		return fmt.Errorf("error removing attribute %s from Account Schema", err)
	}

	d.SetId("")
	return nil
}
