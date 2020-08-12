package main

import (
        "context"
        "fmt"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
        "log"
)

func resourceSource() *schema.Resource {
        return &schema.Resource{
                Create: resourceSourceCreate,
                Read:   resourceSourceRead,
                Update: resourceSourceUpdate,
                Delete: resourceSourceDelete,

                Schema: sourceFields(),
        }
}

func resourceSourceCreate(d *schema.ResourceData, m interface{}) error {
        source, err := expandSource(d)
        if err != nil {
                return err
        }

        log.Printf("[INFO] Creating Source %s", source.Name)

        c, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        newSource, err := c.CreateSource(context.Background(), source)
        if err != nil {
                return err
        }

        err = flattenSource(d, newSource)
        if err != nil {
                return err
        }

        return resourceSourceRead(d, m)
}

func resourceSourceRead(d *schema.ResourceData, m interface{}) error {
        log.Printf("[INFO] Refreshing source ID %s", d.Id())
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

        err = flattenSource(d, source)
        if err != nil {
                return err
        }

        return nil
}

func resourceSourceUpdate(d *schema.ResourceData, m interface{}) error {
        log.Printf("[INFO] Updating Source ID %s", d.Id())
        client, err := m.(*Config).IdentityNowClient()
        if err != nil {
                return err
        }

        updatedSource, err := expandSource(d)
        if err != nil {
                return err
        }

        _, err = client.UpdateSource(context.Background(), updatedSource)
        if err != nil {
                return err
        }

        return resourceSourceRead(d, m)
}

func resourceSourceDelete(d *schema.ResourceData, m interface{}) error {
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
