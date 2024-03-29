package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Flatteners

func flattenSource(d *schema.ResourceData, in *Source) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("connector", in.Connector)
	d.Set("delete_threshold", in.DeleteThreshold)
	d.Set("authoritative", in.Authoritative)
	d.Set("type", in.Type)

	if in.Owner != nil {
		v, ok := d.Get("owner").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("owner", flattenSourceOwner(in.Owner, v))
	}

	if in.Cluster != nil {
		v, ok := d.Get("cluster").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("cluster", flattenSourceCluster(in.Cluster, v))
	}

	if in.AccountCorrelationConfig != nil {
		v, ok := d.Get("account_correlation_config").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("account_correlation_config", flattenSourceAccountCorrelationConfig(in.AccountCorrelationConfig, v))
	}

	if in.ConnectorAttributes != nil {
		v, ok := d.Get("connector_attributes").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("connector_attributes", flattenSourceConnectorAttributes(in.ConnectorAttributes, v))
	}

	if in.Schemas != nil {
		v, ok := d.Get("schemas").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("schemas", flattenSourceSchema(in.Schemas, v))
	}

	if in.ManagementWorkgroup != nil {
		v, ok := d.Get("ManagementWorkgroup").([]interface{})
		if !ok {
			v = []interface{}{}
		}

		d.Set("management_workgroup", flattenSourceManagementWorkgroup(in.ManagementWorkgroup, v))
	}

	if in.PasswordPolicies != nil {
		d.Set("password_policies", flattenSourcePasswordPolicies(in.PasswordPolicies))
	}

	return nil
}

// Expanders

func expandSource(in *schema.ResourceData) (*Source, error) {
	obj := Source{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding source: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)
	obj.Connector = in.Get("connector").(string)
	obj.Type = in.Get("type").(string)

	if v, ok := in.Get("authoritative").(bool); ok {
		obj.Authoritative = v
	}

	if v, ok := in.Get("delete_threshold").(int); ok {
		obj.DeleteThreshold = v
	}

	if v, ok := in.Get("owner").([]interface{}); ok && len(v) > 0 {
		obj.Owner = expandSourceOwner(v)
	}

	if v, ok := in.Get("schemas").([]interface{}); ok && len(v) > 0 {
		obj.Schemas = expandSourceSchema(v)
	}

	if v, ok := in.Get("cluster").([]interface{}); ok && len(v) > 0 {
		obj.Cluster = expandSourceCluster(v)
	}

	if v, ok := in.Get("account_correlation_config").([]interface{}); ok && len(v) > 0 {
		obj.AccountCorrelationConfig = expandSourceAccountCorrelationConfig(v)
	}

	if v, ok := in.Get("connector_attributes").([]interface{}); ok && len(v) > 0 {
		obj.ConnectorAttributes = expandSourceConnectorAttributes(v)
	}

	if v, ok := in.Get("management_workgroup").([]interface{}); ok && len(v) > 0 {
		obj.ManagementWorkgroup = expandSourceManagementWorkgroup(v)
	}

	if v, ok := in.Get("password_policies").([]interface{}); ok && len(v) > 0 {
		obj.PasswordPolicies = expandSourcePasswordPolicies(v)
	}

	return &obj, nil
}
