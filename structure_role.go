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
	d.Set("description", in.Description)
	d.Set("name", in.Name)
	d.Set("requestable", in.Requestable)
	d.Set("enabled", in.Enabled)

	if in.RoleOwner != nil {
		v, ok := d.Get("owner").([]interface{})
		if !ok {
			v = []interface{}{}
		}
		roleOwnerList := []*ObjectInfo{in.RoleOwner}
		d.Set("owner", flattenObjectRoles(roleOwnerList, v))
	}
	if in.AccessProfiles != nil {
		v, ok := d.Get("access_profiles").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("access_profiles", flattenObjectRoles(in.AccessProfiles, v))
	}
	return nil
}

func flattenObjectRoles(in []*ObjectInfo, p []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["type"] = in[i].Type
		obj["id"] = in[i].ID
		obj["name"] = in[i].Name
		out = append(out, obj)
	}
	return out
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

	obj.Description = in.Get("description").(string)
	obj.Name = in.Get("name").(string)

	if v, ok := in.Get("requestable").(bool); ok {
		obj.Requestable = &v
	}

	if v, ok := in.Get("owner").([]interface{}); ok && len(v) > 0 {
		obj.RoleOwner = expandObjectRoles(v)[0]
	}

	if v, ok := in.Get("access_profiles").([]interface{}); ok && len(v) > 0 {
		obj.AccessProfiles = expandObjectRoles(v)
	}

	if v, ok := in.Get("enabled").(bool); ok {
		obj.Enabled = &v
	}

	return &obj, nil
}

func expandUpdateRole(in *schema.ResourceData) ([]*UpdateRole, interface{}, error) {
	updatableFields := []string{"name", "description", "enabled", "owner", "accessProfiles", "requestable"}
	var id interface{}
	if in == nil {
		return nil, nil, fmt.Errorf("[ERROR] Expanding Role: Schema Resource data is nil")
	}

	if v := in.Id(); len(v) > 0 {
		id = v
	}

	out := []*UpdateRole{}

	for i := range updatableFields {
		obj := UpdateRole{}
		if v, ok := in.Get(fmt.Sprintf("/%s", updatableFields[i])).([]interface{}); ok {
			obj.Op = "replace"
			obj.Path = fmt.Sprintf("/%s", updatableFields[i])
			obj.Value = v
		}
		out = append(out, &obj)
	}

	return out, id, nil
}

func expandObjectRoles(p []interface{}) []*ObjectInfo {
	if len(p) == 0 || p[0] == nil {
		return []*ObjectInfo{}
	}
	out := make([]*ObjectInfo, 0, len(p))
	for i := range p {
		obj := ObjectInfo{}
		in := p[i].(map[string]interface{})
		obj.ID = in["id"].(string)
		obj.Name = in["name"].(string)
		obj.Type = in["type"].(string)
		out = append(out, &obj)
	}
	return out
}
