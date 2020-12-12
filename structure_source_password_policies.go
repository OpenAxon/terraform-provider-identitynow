package main

func flattenSourcePasswordPolicies(in []*SourcePasswordPolicies) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["id"] = in[i].ID
		obj["name"] = in[i].Name
		obj["type"] = in[i].Type
		out = append(out, obj)
	}

	return out

}

// Expanders

func expandSourcePasswordPolicies(p []interface{}) []*SourcePasswordPolicies {
	if len(p) == 0 || p[0] == nil {
		return []*SourcePasswordPolicies{}
	}
	out := make([]*SourcePasswordPolicies, 0, len(p))
	for i := range p {
		obj := SourcePasswordPolicies{}
		in := p[i].(map[string]interface{})
		obj.ID = in["id"].(string)
		obj.Name = in["name"].(string)
		obj.Type = in["type"].(string)
		out = append(out, &obj)
	}

	return out
}
