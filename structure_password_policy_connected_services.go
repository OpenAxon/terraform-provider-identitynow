package main

func flattenPasswordPolicyConnectedServices(in []*ConnectedServices) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["id"] = in[i].ID
		obj["external_id"] = in[i].ExternalID
		obj["name"] = in[i].Name
		obj["app_count"] = in[i].AppCount
		obj["supports_password_set_date"] = in[i].SupportsPasswordSetDate
		out = append(out, obj)
	}

	return out
}

func expandPasswordPolicyConnectedServices(p []interface{}) []*ConnectedServices {
	if len(p) == 0 || p[0] == nil {
		return []*ConnectedServices{}
	}
	out := make([]*ConnectedServices, 0, len(p))
	for i := range p {
		obj := ConnectedServices{}
		in := p[i].(map[string]interface{})
		obj.ID = in["id"].(string)
		obj.ExternalID = in["external_id"].(string)
		obj.Name = in["name"].(string)
		obj.AppCount = in["app_count"].(int)
		obj.SupportsPasswordSetDate = in["supports_password_set_date"].(bool)
		out = append(out, &obj)
	}

	return out
}
