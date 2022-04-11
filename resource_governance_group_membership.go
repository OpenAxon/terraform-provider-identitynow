package main

import (
	"context"
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
		log.Printf("member = %s", member)
		membershipRequest.Add = append(membershipRequest.Add, member.(string))
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
		_, notFound := err.(NotFoundError)
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
		_, notFound := err.(NotFoundError)
		if notFound {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	contains := func(g GovernanceGroupMembership, id string) bool {
		for _, member := range g.MemberIDs {
			if member == id {
				return true
			}
		}
		return false
	}
	membershipRequest := GovernanceGroupMembershipRequest{
		Add:    []string{},
		Remove: []string{},
	}
	// First determine what members need to be added if they have been removed from the governance group manually but still defined
	for _, memberFromConfig := range membershipFromConfig.MemberIDs {
		if !contains(*membershipFromRemote, memberFromConfig.(string)) {
			membershipRequest.Add = append(membershipRequest.Add, memberFromConfig.(string))
		}
	}

	// Now determine what members need to be removed if they were added manually to the group but not defined in config
	for _, memberFromRemote := range membershipFromRemote.MemberIDs {
		if !contains(*membershipFromConfig, memberFromRemote.(string)) {
			membershipRequest.Remove = append(membershipRequest.Remove, memberFromRemote.(string))
		}
	}
	err = c.UpdateGovernanceGroupMemberships(context.Background(), membershipFromConfig.GroupID, membershipRequest)
	if err != nil {
		_, notFound := err.(NotFoundError)
		if notFound {
			d.SetId("")
		} else {
			return err
		}
	}
	return nil
}

func resourceGovernanceGroupMembershipDelete(d *schema.ResourceData, m interface{}) error {
	c, err := m.(*Config).IdentityNowClient()
	if err != nil {
		return err
	}

	membership, err := expandGovernanceGroupMembership(d)
	if err != nil {
		return err
	}

	membershipRequest := GovernanceGroupMembershipRequest{
		Remove: []string{},
	}

	for _, memberFromRemote := range membership.MemberIDs {
		membershipRequest.Remove = append(membershipRequest.Remove, memberFromRemote.(string))
	}

	err = c.UpdateGovernanceGroupMemberships(context.Background(), membership.GroupID, membershipRequest)
	if err != nil {
		_, notFound := err.(NotFoundError)
		if notFound {
			d.SetId("")
		} else {
			return err
		}
	}
	return nil
}
