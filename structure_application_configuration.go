package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenApplicationConfiguration(d *schema.ResourceData, in *ApplicationConfiguration) error {
	if d == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("application_id", in.ID)
	d.Set("icon", in.Icon)
	d.Set("account_service_id", in.AccountServiceId)
	d.Set("account_service_match_all_accounts", in.AccountServiceMatchAllAccounts)
	d.Set("access_profile_ids", in.AccessProfileIDs)
	d.Set("app_center_enabled", in.AppCenterEnabled)
	d.Set("launch_pad_enabled", in.LaunchPadEnabled)
	d.Set("provision_request_enabled", in.ProvisionRequestEnabled)
	return nil
}

// Expanders

func expandApplicationConfiguration(in *schema.ResourceData) (*ApplicationConfiguration, error) {
	obj := &ApplicationConfiguration{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Access Profile: Schema Resource data is nil")
	}
	obj.ID = in.Get("application_id").(string)
	obj.Icon = in.Get("icon").(string)
	obj.AccountServiceId = in.Get("account_service_id").(int)
	obj.AccountServiceMatchAllAccounts = in.Get("account_service_match_all_accounts").(bool)
	obj.AccessProfileIDs = in.Get("access_profile_ids").(*schema.Set).List()
	obj.AppCenterEnabled = in.Get("app_center_enabled").(bool)
	obj.LaunchPadEnabled = in.Get("launch_pad_enabled").(bool)
	obj.ProvisionRequestEnabled = in.Get("provision_request_enabled").(bool)
	return obj, nil
}
