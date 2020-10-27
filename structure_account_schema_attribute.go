package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenAccountSchemaAttribute(d *schema.ResourceData, in *AccountSchemaAttribute) error {
	if in == nil {
		return nil
	}

	d.SetId(fmt.Sprintf("%s-%s", in.SourceID, in.Name))
	d.Set("name", in.Name)
	d.Set("type", in.Type)
	d.Set("object_type", in.ObjectType)
	d.Set("source_id", in.SourceID)
	d.Set("description", in.Description)
	d.Set("display_attribute", in.DisplayAttribute)
	d.Set("entitlement", in.Entitlement)
	d.Set("identity_attribute", in.IdentityAttribute)
	d.Set("managed", in.Managed)
	d.Set("minable", in.Minable)
	d.Set("multi", in.Multi)
	return nil
}

// Expanders

func expandAccountSchemaAttribute(in *schema.ResourceData) (*AccountSchemaAttribute, error) {
	obj := AccountSchemaAttribute{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Account Schema Attribute: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Name = in.Get("name").(string)
	obj.Type = in.Get("type").(string)
	obj.ObjectType = in.Get("object_type").(string)
	obj.SourceID = in.Get("source_id").(string)
	obj.Description = in.Get("description").(string)

	if v, ok := in.Get("display_attribute").(bool); ok {
		obj.DisplayAttribute = v
	}

	if v, ok := in.Get("entitlement").(bool); ok {
		obj.Entitlement = v
	}

	if v, ok := in.Get("identity_attribute").(bool); ok {
		obj.IdentityAttribute = v
	}

	if v, ok := in.Get("managed").(bool); ok {
		obj.Managed = v
	}

	if v, ok := in.Get("minable").(bool); ok {
		obj.Minable = v
	}

	if v, ok := in.Get("multi").(bool); ok {
		obj.Multi = v
	}

	return &obj, nil
}

func getAccountSchemaAttribute(accountSchema *AccountSchema, name string) *AccountSchemaAttribute {
	attributes := accountSchema.Attributes
	for i := range attributes {
		if attributes[i].Name == name {
			return attributes[i]
		}
	}
	return nil
}
