package model

// TenantWithGroups model
type TenantWithGroups struct {
	// Id of Tenant
	Id *string `json:"id,omitempty"`
	// Email of Tenant
	Email            *string          `json:"email,omitempty"`
	InvitationStatus InvitationStatus `json:"invitationStatus,omitempty"`
	// List of all groups of Tenant
	Groups []TenantGroupDetail `json:"groups,omitempty"`
}
