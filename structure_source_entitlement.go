package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenSourceEntitlement(d *schema.ResourceData, in *Items) error {
	if in == nil {
		return nil
	}
	d.SetId(in.ID)
	d.Set("source_id", in.SourceID)
	d.Set("source_name", in.SourceName)
	d.Set("attribute", in.Attribute)
	d.Set("created_time", in.CreatedTime)
	d.Set("deleted_time", in.DeletedTime)
	d.Set("description", in.Description)
	d.Set("direct_permissions", toArrayString(in.DirectPermissions))
	d.Set("display_name", in.DisplayName)
	d.Set("displayable_name", in.DisplayableName)
	d.Set("last_modified_time", in.LastModifiedTime)
	d.Set("owner_id", in.OwnerID)
	d.Set("owner_uid", in.OwnerUID)
	d.Set("privileged", in.Privileged)
	d.Set("schema", in.Schema)
	d.Set("value", in.Value)

	return nil
}

func getEntitlement(entitlements []*Items, name string) (*Items, error) {
	for i := range entitlements {
		if entitlements[i].DisplayableName == name {
			return entitlements[i], nil
		}
	}
	return nil, NotFoundError{fmt.Sprintf("no entitlement named '%s' could be found", name)}
}
