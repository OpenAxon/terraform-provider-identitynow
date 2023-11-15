package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceAccountSchemaImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	sourceID, schemaId, err := splitAccountSchemaID(d.Id())
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	d.Set("source_id", sourceID)
	d.Set("schema_id", schemaId)
	err = resourceAccountSchemaRead(d, meta)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
