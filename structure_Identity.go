package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func flattenIdentity(d *schema.ResourceData, in *Identity) error {
	if in == nil {
		return nil
	}
	d.SetId(in.ID)
	d.Set("alias", in.Alias)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("enabled", in.Enabled)
	d.Set("isManager", in.IsManager)
	d.Set("emailAddress", in.EmailAddress)
	d.Set("identityStatus", in.IdentityStatus)

	if in.IdentityAttributes != nil {
		v, ok := d.Get("attributes").(interface{})
		if !ok {
			v = []interface{}{}
		}
		d.Set("attributes", flattenIdentityAttributes(in.IdentityAttributes, v))
	}

	return nil
}

func flattenIdentityAttributes(in *IdentityAttributes, p interface{}) interface{} {
	if in == nil {
		return []interface{}{}
	}
	var obj = make(map[string]interface{})
	obj["adpId"] = in.AdpID
	obj["lastname"] = in.LastName
	obj["firstname"] = in.FirstName
	obj["phone"] = in.Phone
	obj["userType"] = in.UserType
	obj["uid"] = in.UID
	obj["email"] = in.Email
	obj["workdayId"] = in.WorkdayId

	return obj
}
