package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
        providerDefaultEmptyString = "nil"
)

var (
        descriptions                        map[string]string
)

func Provider() *schema.Provider {
        return &schema.Provider{
                Schema: map[string]*schema.Schema{
                        "api_url": {
                                Type:        schema.TypeString,
                                Required:    true,
                                DefaultFunc: schema.EnvDefaultFunc("IDENTITYNOW_URL", providerDefaultEmptyString),
                                Description: descriptions["api_url"],
                        },
                        "client_id": {
                                Type:        schema.TypeString,
                                Optional:    true,
                                Sensitive:   true,
                                DefaultFunc: schema.EnvDefaultFunc("IDENTITYNOW_CLIENT_ID", providerDefaultEmptyString),
                                Description: descriptions["client_id"],
                        },
                        "client_secret": {
                                Type:        schema.TypeString,
                                Optional:    true,
                                Sensitive:   true,
                                DefaultFunc: schema.EnvDefaultFunc("IDENTITYNOW_CLIENT_SECRET", providerDefaultEmptyString),
                                Description: descriptions["client_secret"],
                        },
                },

                ResourcesMap: map[string]*schema.Resource{
	                "identitynow_source_azure_ad" : resourceSourceAAD(),
                },

                DataSourcesMap: map[string]*schema.Resource{
                    "identitynow_source_azure_ad" : dataSourceSource(),
                },

                ConfigureFunc: providerConfigure,
        }
}

func init() {
        descriptions = map[string]string{
                "api_url":    "The URL to the IdentityNow API",
                "client_id":  "API client used to authenticate with the IdentityNow API",
                "client_secret": "API client secret used to authenticate with the IdentityNow API",
        }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
        apiURL := d.Get("api_url").(string)
        clientId := d.Get("client_id").(string)
        clientSecret := d.Get("client_secret").(string)

        config := &Config{
                URL:       apiURL,
                ClientId:  clientId,
                ClientSecret:   clientSecret,
        }

        return config, nil
}
