package model

// OrganizationPendingInvitation model
type OrganizationPendingInvitation struct {
	// Id of Tenant (required)
	Id *string `json:"id,omitempty"`
	// Email of Tenant (required)
	Email *string `json:"email,omitempty"`
	// Id of group (required)
	GroupId *string `json:"groupId,omitempty"`
	// Name of group (required)
	GroupName *string `json:"groupName,omitempty"`
}
