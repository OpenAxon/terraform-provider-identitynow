package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSourceAAD() *schema.Resource {
        return &schema.Resource{
                Create: resourceSourceAADCreate,
                Read:   resourceSourceAADRead,
                Update: resourceSourceAADUpdate,
                Delete: resourceSourceAADDelete,

                Schema: map[string]*schema.Schema{
                        "ms_graph_token_base": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceSourceAADCreate(d *schema.ResourceData, m interface{}) error {
        //msGraphTokenBase := d.Get("ms_graph_token_base").(string)
        d.SetId("source_aad_1234")
        return resourceSourceAADRead(d, m)
}

func resourceSourceAADRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceSourceAADUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceSourceAADRead(d, m)
}

func resourceSourceAADDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
