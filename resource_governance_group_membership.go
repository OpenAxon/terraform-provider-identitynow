package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGovernanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGovernanceGroupMembershipCreate,
		Read:   resourceGovernanceGroupMembershipRead,
		Update: resourceGovernanceGroupMembershipUpdate,
		Delete: resourceGovernanceGroupMembershipDelete,
		Schema: governanceGroupMembershipFields(),
	}
}

func resourceGovernanceGroupMembershipCreate(d *schema.ResourceData, m interface{}) error {
	membership, err := expandGovernanceGroupMembership(d)
	if err != nil {
		return err
	}

	c, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	membershipRequest := GovernanceGroupMembershipRequest{
		Add:    []string{},
		Remove: []string{},
	}
	for _, member := range membership.MemberIDs {
		membershipRequest.Add = append(membershipRequest.Add, member)
	}

	err = c.UpdateGovernanceGroupMemberships(context.Background(), membership.GroupID, membershipRequest)
	if err != nil {
		return err
	}

	err = flattenGovernanceGroupMembership(d, membership)
	if err != nil {
		return err
	}
	return resourceGovernanceGroupMembershipRead(d, m)
}

func resourceGovernanceGroupMembershipRead(d *schema.ResourceData, m interface{}) error {
	client, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	membership, err := client.GetGovernanceGroupMembership(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	err = flattenGovernanceGroupMembership(d, membership)
	if err != nil {
		return err
	}

	return nil
}

func resourceGovernanceGroupMembershipUpdate(d *schema.ResourceData, m interface{}) error {
	c, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	membershipFromConfig, err := expandGovernanceGroupMembership(d)
	if err != nil {
		return err
	}

	membershipFromRemote, err := c.GetGovernanceGroupMembership(context.Background(), d.Id())
	if err != nil {
		_, notFound := err.(*NotFoundError)
		if notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	contains := func(g GovernanceGroupMembership, id string) bool {
		for _, member := range g.MemberIDs {
			log.Printf("[DEBUG] value = %+v", member)
			if member == id {
				return true
			}
		}
		return false
	}

	// First determine what members need to be added if they have been removed from the governance group manually but still defined
	toBeAdded := []string{}
	for _, memberFromConfig := range membershipFromConfig.MemberIDs {
		if !contains(*membershipFromRemote, memberFromConfig) {
			toBeAdded = append(toBeAdded, memberFromConfig)
		}
	}

	// Now determine what members need to be removed if they were added manually to the group but not defined in config
	toBeRemoved := []string{}
	for _, memberFromRemote := range membershipFromRemote.MemberIDs {
		if !contains(*membershipFromConfig, memberFromRemote) {
			toBeRemoved = append(toBeRemoved, memberFromRemote)
		}
	}

	membershipRequest := GovernanceGroupMembershipRequest{
		Add:    toBeAdded,
		Remove: toBeRemoved,
	}

	err = c.UpdateGovernanceGroupMemberships(context.Background(), membershipFromConfig.GroupID, membershipRequest)
	if err != nil {
		return err
	}
	return nil
}

func resourceGovernanceGroupMembershipDelete(d *schema.ResourceData, m interface{}) error {
	return fmt.Errorf("resourceGovernanceGroupMembershipDelete not implemented")

	// log.Printf("[INFO] Deleting Governance Group ID %s", d.Id())

	// client, err := m.(*Config).IdentityNowClient()
	// if err != nil {
	// 	return err
	// }

	// governanceGroup, err := client.GetGovernanceGroup(context.Background(), d.Id())
	// if err != nil {
	// 	_, notFound := err.(*NotFoundError)
	// 	if notFound {
	// 		log.Printf("[INFO] Governance Group ID %s not found.", d.Id())
	// 		d.SetId("")
	// 		return nil
	// 	}
	// 	return err
	// }

	// res, err := client.DeleteGovernanceGroup(context.Background(), governanceGroup.ID)
	// if err != nil {
	// 	return err
	// }

	// if len(res.Deleted) != 1 {
	// 	return fmt.Errorf("expected result id array to be 1, got %d :%v", len(res.Deleted), res.Deleted)
	// }

	// if res.Deleted[0] != governanceGroup.ID {
	// 	return fmt.Errorf("expected result id to be %s, got %s", governanceGroup.ID, res.Deleted[0])
	// }

	// d.SetId("")
	// return nil
}
