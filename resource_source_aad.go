package main

import (
        "context"
        "fmt"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
        "log"
)

func resourceSourceAAD() *schema.Resource {
        return &schema.Resource{
                Create: resourceSourceAADCreate,
                Read:   resourceSourceAADRead,
                Update: resourceSourceAADUpdate,
                Delete: resourceSourceAADDelete,

                Schema: sourceFields(),
        }
}

func resourceSourceAADCreate(d *schema.ResourceData, m interface{}) error {
        //msGraphTokenBase := d.Get("ms_graph_token_base").(string)
        c := NewClient("c70cde50e14d4e5e9082392056f9faf3", "e622774e2d52c6e6d5c5f2c7bea5134e9a892adc63c2529c04467064db9b2ba1")
        ctx := context.Background()
        //res, err := c.GetSource(ctx, "2c9180887349a97a01734a1e6441069f")
        if err := c.GetToken(ctx); err != nil {
                return err
        }
        if len(c.accessToken) == 0 {
                return fmt.Errorf("access token is empty")
        }
        log.Printf(c.accessToken)

        source, err := c.GetSource(ctx, "2c9180887349a97a01734a1e6441069f")
        if err != nil {
                return err
        }

        err = flattenSourceAAD(d, source)
        if err != nil {
                return err
        }
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
