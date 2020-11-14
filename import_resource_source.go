package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceSourceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	err := resourceSourceRead(d, meta)

	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
