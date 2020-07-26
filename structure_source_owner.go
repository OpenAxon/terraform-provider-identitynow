package main

// Flatteners

func flattenSourceOwner(in *Owner, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	obj["type"] = in.Type
	obj["id"] = in.ID
	obj["name"] = in.Name

	return []interface{}{obj}

}
