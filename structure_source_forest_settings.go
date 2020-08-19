package main

// Flatteners

func flattenSourceForestSettings(in []*ForestSettings, p []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["password"] = in[i].Password
		obj["gc_server"] = in[i].GcServer
		obj["forest_name"] = in[i].ForestName
		obj["user"] = in[i].User
		obj["use_ssl"] = in[i].UseSSL
		obj["authorization_type"] = in[i].AuthorizationType
		out = append(out, obj)
	}

	return out

}

// Expanders

func expandSourceForestSettings(p []interface{}) []*ForestSettings {
	if len(p) == 0 || p[0] == nil {
		return []*ForestSettings{}
	}
	out := make([]*ForestSettings, 0, len(p))

	for i := range p {
		obj := ForestSettings{}
		in := p[i].(map[string]interface{})
		obj.Password = in["password"].(string)
		obj.GcServer = in["gc_server"].(string)
		obj.ForestName = in["forest_name"].(string)
		obj.User = in["user"].(string)
		obj.UseSSL = in["use_ssl"].(bool)
		obj.AuthorizationType = in["authorization_type"].(string)
		out = append(out, &obj)
	}

	return out
}
