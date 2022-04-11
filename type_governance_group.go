package main

type GovernanceGroupOwner struct {
	ID string `json:"id"`
}

type GovernanceGroup struct {
	ID             string               `json:"id,omitempty"`
	Name           string               `json:"name,omitempty"`
	Description    string               `json:"description,omitempty"`
	Owner          GovernanceGroupOwner `json:"owner,omitempty"`
	ApprovalScheme string
}

type GovernanceGroupBulkDeletionRequest struct {
	IDs []string `json:"ids,omitempty"`
}

type GovernanceGroupBulkDeletionResponse struct {
	Deleted []string `json:"deleted,omitempty"`
}
