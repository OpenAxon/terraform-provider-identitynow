package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenAccessProfile(d *schema.ResourceData, in *AccessProfile) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("source_id", in.SourceID)
	d.Set("source_name", in.SourceName)
	d.Set("owner_id", in.OwnerID)
	d.Set("entitlements", toArrayInterface(in.Entitlements))
	d.Set("denied_comments_required", in.DeniedCommentsRequired)
	d.Set("disabled", in.Disabled)
	d.Set("protected", in.Protected)
	d.Set("request_comments_required", in.RequestCommentsRequired)
	d.Set("requestable", in.Requestable)
	d.Set("approval_schemes", in.ApprovalSchemes)
	d.Set("revoke_request_approval_schemes", in.RevokeRequestApprovalSchemes)
	return nil
}

// Expanders

func expandAccessProfile(in *schema.ResourceData) (*AccessProfile, error) {
	obj := AccessProfile{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Access Profile: Schema Resource data is nil")
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)
	obj.ApprovalSchemes = in.Get("approval_schemes").(string)
	obj.RevokeRequestApprovalSchemes = in.Get("revoke_request_approval_schemes").(string)

	if v, ok := in.Get("source_id").(int); ok {
		obj.SourceID = v
	}

	if v, ok := in.Get("entitlements").([]interface{}); ok {
		obj.Entitlements = toArrayString(v)
	}

	if v, ok := in.Get("owner_id").(int); ok {
		obj.OwnerID = v
	}

	if v, ok := in.Get("denied_comments_required").(bool); ok {
		obj.DeniedCommentsRequired = &v
	}

	if v, ok := in.Get("disabled").(bool); ok {
		obj.Disabled = &v
	}

	if v, ok := in.Get("protected").(bool); ok {
		obj.Protected = &v
	}

	if v, ok := in.Get("request_comments_required").(bool); ok {
		obj.RequestCommentsRequired = &v
	}

	return &obj, nil
}
