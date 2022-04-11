package main

type GovernanceGroupMembership struct {
	GroupID   string   `json:"group_id,omitempty"`
	MemberIDs []string `json:"member_ids,omitempty"`
}

type GovernanceGroupMembershipRequest struct {
	Add    []string `json:"add"`
	Remove []string `json:"remove"`
}

type GovernanceGroupMember struct {
	ExternalID string `json:"externalId"`
}
