package main

// Flatteners

func flattenSourceSearchDNs(in []*SearchDNs) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))

	for i := range in {
		var obj = make(map[string]interface{})
		obj["group_membership_search_dn"] = in[i].GroupMembershipSearchDN
		obj["search_dn"] = in[i].SearchDN
		obj["group_member_filter_string"] = in[i].GroupMemberFilterString
		obj["search_scope"] = in[i].SearchScope
		obj["primary_group_search_dn"] = in[i].PrimaryGroupSearchDN
		obj["iterate_search_filter"] = in[i].IterateSearchFilter
		out = append(out, obj)
	}

	return out

}

// Expanders

func expandSourceSearchDNs(p []interface{}) []*SearchDNs {
	if len(p) == 0 || p[0] == nil {
		return []*SearchDNs{}
	}
	out := make([]*SearchDNs, 0, len(p))
	for i := range p {
		obj := SearchDNs{}
		in := p[i].(map[string]interface{})
		obj.GroupMembershipSearchDN = in["group_membership_search_dn"].(string)
		obj.SearchDN = in["search_dn"].(string)
		obj.GroupMemberFilterString = in["group_member_filter_string"].(string)
		obj.SearchScope = in["search_scope"].(string)
		obj.PrimaryGroupSearchDN = in["primary_group_search_dn"].(string)
		obj.IterateSearchFilter = in["iterate_search_filter"].(string)
		out = append(out, &obj)
	}

	return out
}
