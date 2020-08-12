package main

// Flatteners

func flattenSourceSchema(in []*Schema, p []interface{}) []interface{} {
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

func expandSourceSchema(p []interface{}) []*Schema {
	if len(p) == 0 || p[0] == nil {
		return []*Schema{}
	}
	out := make([]*Schema, 0, len(p))
	for i := range p {
		obj := Schema{}
		in := p[i].(map[string]interface{})
		obj.ID = in["id"].(string)
		obj.Name = in["name"].(string)
		obj.Type = in["type"].(string)
		out = append(out, &obj)
	}

	return out
}