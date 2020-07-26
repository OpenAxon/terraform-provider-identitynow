package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenSourceAAD(d *schema.ResourceData, in *SourceAAD) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("delete_threshold", in.DeleteThreshold)
	d.Set("authoritative", in.Authoritative)

	if in.Owner != nil {
		v, ok := d.Get("owner").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("owner", flattenSourceOwner(in.Owner, v))
	}

	return nil
}

// Expanders

func expandSourceAAD(in *schema.ResourceData) (*SourceAAD, error) {
	obj := SourceAAD{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding source: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)
	obj.Connector = "azure-active-directory"

	if v, ok := in.Get("authoritative").(bool); ok {
		obj.Authoritative = v
	}

	if v, ok := in.Get("delete_threshold").(int); ok {
		obj.DeleteThreshold = v
	}

	if v, ok := in.Get("owner").([]interface{}); ok && len(v) > 0 {
		obj.Owner = expandSourceOwner(v)
	}

	return &obj, nil
}