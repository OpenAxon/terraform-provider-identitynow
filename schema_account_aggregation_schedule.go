package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func accountAggregationScheduleFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"source_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source ID",
		},

		"cron_expressions": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Account aggregation scheduling in cron Expression format.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	return s
}
