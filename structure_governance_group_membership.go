package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenGovernanceGroupMembership(d *schema.ResourceData, in *GovernanceGroupMembership) error {
	if in == nil {
		return nil
	}
	d.SetId(in.GroupID)
	d.Set("group_id", in.GroupID)
	d.Set("member_ids", in.MemberIDs)
	return nil
}

func expandGovernanceGroupMembership(in *schema.ResourceData) (*GovernanceGroupMembership, error) {
	obj := GovernanceGroupMembership{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Governance Group: Schema Resource data is nil")
	}

	obj.GroupID = in.Get("group_id").(string)
	obj.MemberIDs = in.Get("member_ids").([]string)
	return &obj, nil
}
