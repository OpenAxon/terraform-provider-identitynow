package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testAccountAggregationScheduleConf      *AccountAggregationSchedule
	testAccountAggregationScheduleInterface map[string]interface{}
)

func init() {
	testAccountAggregationScheduleConf = &AccountAggregationSchedule{
		SourceID:        "1234",
		CronExpressions: []string{"0 0 1,2,3,4 * * ?"},
	}
	testAccountAggregationScheduleInterface = map[string]interface{}{
		"source_id":        "1234",
		"cron_expressions": []interface{}{"0 0 1,2,3,4 * * ?"},
	}
}

func TestFlattenAccountAggregationSchedule(t *testing.T) {
	cases := []struct {
		Input          *AccountAggregationSchedule
		ExpectedOutput map[string]interface{}
	}{
		{
			testAccountAggregationScheduleConf,
			testAccountAggregationScheduleInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, accountAggregationScheduleFields(), tc.ExpectedOutput)
		err := flattenAccountAggregationSchedule(output, tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(expectedOutput, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, expectedOutput)
		}
	}
}

func TestExpandAccountAggregationSchedule(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *AccountAggregationSchedule
	}{
		{
			testAccountAggregationScheduleInterface,
			testAccountAggregationScheduleConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, accountAggregationScheduleFields(), tc.Input)
		output, err := expandAccountAggregationSchedule(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
