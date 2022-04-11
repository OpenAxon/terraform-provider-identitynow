package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenGovernanceGroup(d *schema.ResourceData, in *GovernanceGroup) error {
	if in == nil {
		return nil
	}
	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("owner_id", in.Owner.ID)
	return nil
}

func expandGovernanceGroup(in *schema.ResourceData) (*GovernanceGroup, error) {
	obj := GovernanceGroup{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Governance Group: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)
	obj.Owner.ID = in.Get("owner_id").(string)
	return &obj, nil
}
