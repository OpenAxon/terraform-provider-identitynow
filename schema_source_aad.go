package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func sourceFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source name",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source description",
		},
		"delete_threshold": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      10,
		},
		"authoritative": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "True if this source is authoritative",
			Default:     false,
		},
		"owner": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Required:      true,
			Elem: &schema.Resource{
				Schema: sourceOwnerFields(),
			},
		},
		"schemas": {
			Type:          schema.TypeList,
			Computed:      true,
			Elem: &schema.Resource{
				Schema: sourceSchemaFields(),
			},
		},
		"cluster": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			Elem: &schema.Resource{
				Schema: sourceClusterFields(),
			},
		},
		"account_correlation_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Computed:      true,
			Elem: &schema.Resource{
				Schema: sourceAccountCorrelationConfigFields(),
			},
		},
		"connector_attributes": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			Elem: &schema.Resource{
				Schema: sourceConnectorAttributesFields(),
			},
		},
	}

	//for k, v := range commonAnnotationLabelFields() {
	//	s[k] = v
	//}

	return s
}