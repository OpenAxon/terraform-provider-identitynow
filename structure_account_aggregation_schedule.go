package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenAccountAggregationSchedule(d *schema.ResourceData, in *AccountAggregationSchedule) error {
	if in == nil {
		return nil
	}

	d.SetId(in.SourceID)
	d.Set("cron_expressions", toArrayInterface(in.CronExpressions))
	return nil
}

// Expanders

func expandAccountAggregationSchedule(in *schema.ResourceData) (*AccountAggregationSchedule, error) {
	obj := AccountAggregationSchedule{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Schedule Account Aggregation: Schema Resource data is nil")
	}

	obj.SourceID = in.Get("source_id").(string)

	if v, ok := in.Get("cron_expressions").([]interface{}); ok && len(v) > 0 {
		obj.CronExpressions = toArrayString(v)
	}

	return &obj, nil
}
