package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenApplication(d *schema.ResourceData, in *Application) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("app_id", in.AppID)
	return nil
}

// Expanders

func expandApplication(in *schema.ResourceData) (*Application, error) {
	obj := &Application{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Access Profile: Schema Resource data is nil")
	}
	obj.ID = in.Id()
	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)
	obj.AppID = in.Get("app_id").(string)
	return obj, nil
}
