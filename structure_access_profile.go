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
	d.Set("enabled", in.Enabled)
	d.Set("requestable", in.Requestable)

	if in.AccessProfileOwner != nil {
		v, ok := d.Get("owner").([]interface{})
		if !ok {
			v = []interface{}{}
		}
		accessProfileOwnerList := []*ObjectInfo{in.AccessProfileOwner}
		d.Set("owner", flattenObjectAccessProfile(accessProfileOwnerList, v))
	}
	if in.AccessProfileSource != nil {
		v, ok := d.Get("source").([]interface{})
		if !ok {
			v = []interface{}{}
		}
		accessProfileSourceList := []*ObjectInfo{in.AccessProfileSource}
		d.Set("source", flattenObjectAccessProfile(accessProfileSourceList, v))
	}
	if in.Entitlements != nil {
		v, ok := d.Get("entitlements").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("entitlements", flattenObjectRoles(in.Entitlements, v))
	}
	return nil
}

func flattenObjectAccessProfile(in []*ObjectInfo, p []interface{}) []interface{} {
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

func expandAccessProfile(in *schema.ResourceData) (*AccessProfile, error) {
	obj := AccessProfile{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Access Profile: Schema Resource data is nil")
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)

	if v, ok := in.Get("enabled").(bool); ok {
		obj.Enabled = &v
	}

	if v, ok := in.Get("owner").([]interface{}); ok && len(v) > 0 {
		obj.AccessProfileOwner = expandObjectAccessProfile(v)[0]
	}

	if v, ok := in.Get("source").([]interface{}); ok && len(v) > 0 {
		obj.AccessProfileSource = expandObjectAccessProfile(v)[0]
	}

	if v, ok := in.Get("entitlements").([]interface{}); ok && len(v) > 0 {
		obj.Entitlements = expandObjectRoles(v)
	}

	return &obj, nil
}

func expandUpdateAccessProfile(in *schema.ResourceData) ([]*UpdateAccessProfile, interface{}, error) {
	updatableFields := []string{"name", "description", "enabled", "owner", "entitlements", "requestable", "source"}
	var id interface{}
	if in == nil {
		return nil, nil, fmt.Errorf("[ERROR] Expanding Role: Schema Resource data is nil")
	}

	if v := in.Id(); len(v) > 0 {
		id = v
	}

	out := []*UpdateAccessProfile{}

	for i := range updatableFields {
		obj := UpdateAccessProfile{}
		if v, ok := in.Get(fmt.Sprintf("/%s", updatableFields[i])).([]interface{}); ok {
			obj.Op = "replace"
			obj.Path = fmt.Sprintf("/%s", updatableFields[i])
			obj.Value = v
		}
		out = append(out, &obj)
	}

	return out, id, nil
}

func expandObjectAccessProfile(p []interface{}) []*ObjectInfo {
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
