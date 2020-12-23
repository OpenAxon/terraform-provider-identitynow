package main

func flattenSourceGroupSearchDNs(in []*GroupSearchDNs) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["search_dn"] = in[i].SearchDN
		obj["search_scope"] = in[i].SearchScope
		obj["iterate_search_filter"] = in[i].IterateSearchFilter
		out = append(out, obj)
	}

	return out

}

// Expanders

func expandSourceGroupSearchDNs(p []interface{}) []*GroupSearchDNs {
	if len(p) == 0 || p[0] == nil {
		return []*GroupSearchDNs{}
	}
	out := make([]*GroupSearchDNs, 0, len(p))
	for i := range p {
		obj := GroupSearchDNs{}
		in := p[i].(map[string]interface{})
		obj.SearchDN = in["search_dn"].(string)
		obj.SearchScope = in["search_scope"].(string)
		obj.IterateSearchFilter = in["iterate_search_filter"].(string)
		out = append(out, &obj)
	}

	return out
}
