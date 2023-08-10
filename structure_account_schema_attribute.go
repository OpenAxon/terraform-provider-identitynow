package main

// Flatteners

func flattenAccountSchemaAttributes(in []*AccountSchemaAttribute, p []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))

	for i := range in {
		var obj = make(map[string]interface{})
		obj["name"] = in[i].Name
		obj["type"] = in[i].Type
		obj["description"] = in[i].Description
		obj["is_multi_valued"] = in[i].IsMultiValued
		obj["is_entitlement"] = in[i].IsEntitlement
		obj["is_group"] = in[i].IsGroup
		if in[i].Schema != nil {
			v, ok := obj["schema"].([]interface{})
			if !ok {
				v = []interface{}{}
			}
			obj["schema"] = flattenAccountSchemaAttributesSchema(in[i].Schema, v)
		}
		out = append(out, obj)
	}
	return out
}

func flattenAccountSchemaAttributesSchema(in *AccountSchemaAttributeSchema, p []interface{}) interface{} {
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

// Expanders
func expandAccountSchemaAttributes(p []interface{}) []*AccountSchemaAttribute {
	if len(p) == 0 || p[0] == nil {
		return []*AccountSchemaAttribute{}
	}
	out := make([]*AccountSchemaAttribute, 0, len(p))
	for i := range p {
		obj := AccountSchemaAttribute{}
		in := p[i].(map[string]interface{})
		obj.Name = in["name"].(string)
		obj.Type = in["type"].(string)
		if v, ok := in["description"].(string); ok {
			obj.Description = v
		}

		if v, ok := in["is_multi_valued"].(bool); ok {
			obj.IsMultiValued = v
		}
		if v, ok := in["is_entitlement"].(bool); ok {
			obj.IsEntitlement = v
		}

		if v, ok := in["is_group"].(bool); ok {
			obj.IsGroup = v
		}
		if v, ok := in["schema"].([]interface{}); ok && len(v) > 0 {
			obj.Schema = expandAccountSchemaAttributesSchema(v)
		}
		out = append(out, &obj)
	}

	return out
}

func expandAccountSchemaAttributesSchema(p []interface{}) *AccountSchemaAttributeSchema {
	obj := AccountSchemaAttributeSchema{}

	if len(p) == 0 || p[0] == nil {
		return &obj
	}
	in := p[0].(map[string]interface{})

	obj.ID = in["id"].(string)
	obj.Name = in["name"].(string)
	obj.Type = in["type"].(string)

	return &obj
}

func getAccountSchemaAttribute(accountSchema *AccountSchema, name string) *AccountSchemaAttribute {
	attributes := accountSchema.Attributes
	for i := range attributes {
		if attributes[i].Name == name {
			return attributes[i]
		}
	}
	return nil
}
