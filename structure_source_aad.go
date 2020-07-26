package main

import (
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