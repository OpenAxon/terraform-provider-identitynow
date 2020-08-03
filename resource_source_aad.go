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
        sourceAAD, err := expandSourceAAD(d)
        if err != nil {
                return err
        }

        log.Printf("[INFO] Creating Source AAD %s", sourceAAD.Name)

        c, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        newSource, err := c.CreateSource(context.Background(), sourceAAD)
        if err != nil {
                return err
        }

        err = flattenSourceAAD(d, newSource)
        if err != nil {
                return err
        }

        return resourceSourceAADRead(d, m)
}

func resourceSourceAADRead(d *schema.ResourceData, m interface{}) error {
        log.Printf("[INFO] Refreshing AAD source ID %s", d.Id())
        client, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        source, err := client.GetSource(context.Background(), d.Id())
        if err != nil {
                // non-panicking type assertion, 2nd arg is boolean indicating type match
                _, notFound := err.(*NotFoundError)
                if notFound {
                        log.Printf("[INFO] Source ID %s not found.", d.Id())
                        d.SetId("")
                        return nil
                }
                return err
        }

        err = flattenSourceAAD(d, source)
        if err != nil {
                return err
        }

        return nil
}

func resourceSourceAADUpdate(d *schema.ResourceData, m interface{}) error {
        log.Printf("[INFO] Updating AAD Source ID %s", d.Id())
        client, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        updatedSource, err := expandSourceAAD(d)
        if err != nil {
                return err
        }

        _, err = client.UpdateSource(context.Background(), updatedSource)
        if err != nil {
                return err
        }

        return resourceSourceAADRead(d, m)
}

func resourceSourceAADDelete(d *schema.ResourceData, m interface{}) error {
        log.Printf("[INFO] Deleting Source ID %s", d.Id())

        client, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        source, err := client.GetSource(context.Background(), d.Id())
        if err != nil {
                // non-panicking type assertion, 2nd arg is boolean indicating type match
                _, notFound := err.(*NotFoundError)
                if notFound {
                        log.Printf("[INFO] Source ID %s not found.", d.Id())
                        d.SetId("")
                        return nil
                }
                return err
        }

        err = client.DeleteSource(context.Background(), source)
        if err != nil {
                return fmt.Errorf("Error removing Source: %s", err)
        }

        d.SetId("")
        return nil
}
