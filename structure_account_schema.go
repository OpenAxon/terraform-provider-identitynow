package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenAccountSchema(d *schema.ResourceData, in *AccountSchema) error {
	if in == nil {
		return nil
	}

	d.Set("name", in.ID)
	d.Set("name", in.Name)
	d.Set("source_id", in.SourceID)
	d.Set("schema_id", in.ID)
	d.Set("native_object_type", in.NativeObjectType)
	d.Set("identity_attribute", in.IdentityAttribute)
	d.Set("display_attribute", in.DisplayAttribute)
	d.Set("hierarchy_attribute", in.HierarchyAttribute)
	d.Set("include_permissions", in.IncludePermissions)
	d.Set("modified", in.Modified)
	d.Set("created", in.Created)
	if in.Attributes != nil {
		v, ok := d.Get("attributes").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("attributes", flattenAccountSchemaAttributes(in.Attributes, v))
	}
	return nil
}

// Expanders
func expandAccountSchema(in *schema.ResourceData) (*AccountSchema, error) {
	obj := AccountSchema{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Account Schema: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}
	obj.Name = in.Get("name").(string)
	obj.SourceID = in.Get("source_id").(string)
	obj.ID = in.Get("schema_id").(string)
	obj.NativeObjectType = in.Get("native_object_type").(string)
	obj.IdentityAttribute = in.Get("identity_attribute").(string)
	obj.DisplayAttribute = in.Get("display_attribute").(string)
	obj.HierarchyAttribute = in.Get("hierarchy_attribute").(string)
	obj.Modified = in.Get("modified").(string)
	obj.Created = in.Get("created").(string)
	if v, ok := in.Get("include_permissions").(bool); ok {
		obj.IncludePermissions = v
	}
	if v, ok := in.Get("attributes").([]interface{}); ok && len(v) > 0 {
		obj.Attributes = expandAccountSchemaAttributes(v)
	}
	return &obj, nil
}
