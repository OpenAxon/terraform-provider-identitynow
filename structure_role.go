package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenRole(d *schema.ResourceData, in *Role) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("access_profile_ids", toArrayInterface(in.AccessProfileIds))
	d.Set("approval_schemes", in.ApprovalSchemes)
	d.Set("denied_comments_required", in.DeniedCommentsRequired)
	d.Set("description", in.Description)
	d.Set("disabled", in.Disabled)
	d.Set("display_name", in.DisplayName)
	d.Set("identity_count", in.IdentityCount)
	d.Set("name", in.Name)
	d.Set("owner", in.Owner)
	d.Set("request_comments_required", in.RequestCommentsRequired)
	d.Set("requestable", in.Requestable)
	d.Set("revoke_request_approval_schemes", in.RevokeRequestApprovalSchemes)
	return nil
}

// Expanders

func expandRole(in *schema.ResourceData) (*Role, error) {
	obj := Role{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Role: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.ApprovalSchemes = in.Get("approval_schemes").(string)
	obj.Description = in.Get("description").(string)
	obj.DisplayName = in.Get("display_name").(string)
	obj.Name = in.Get("name").(string)
	obj.Owner = in.Get("owner").(string)
	obj.RevokeRequestApprovalSchemes = in.Get("revoke_request_approval_schemes").(string)

	if v, ok := in.Get("access_profile_ids").([]interface{}); ok {
		obj.AccessProfileIds = toArrayString(v)
	}

	if v, ok := in.Get("identity_count").(int); ok {
		obj.IdentityCount = v
	}

	if v, ok := in.Get("denied_comments_required").(bool); ok {
		obj.DeniedCommentsRequired = &v
	}

	if v, ok := in.Get("disabled").(bool); ok {
		obj.Disabled = &v
	}

	if v, ok := in.Get("request_comments_required").(bool); ok {
		obj.RequestCommentsRequired = &v
	}

	if v, ok := in.Get("requestable").(bool); ok {
		obj.Requestable = &v
	}

	return &obj, nil
}
