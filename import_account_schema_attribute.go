package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceAccountSchemaAttributeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	sourceID, name, err := splitAccountSchemaAttributeID(d.Id())
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	d.Set("source_id", sourceID)
	d.Set("name", name)
	err = resourceAccountSchemaAttributeRead(d, meta)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
