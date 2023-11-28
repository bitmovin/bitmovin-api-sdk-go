package model

// TenantWithGroups model
type TenantWithGroups struct {
	// Id of Tenant (required)
	Id *string `json:"id,omitempty"`
	// Email of Tenant (required)
	Email *string `json:"email,omitempty"`
	// List of all groups of Tenant (required)
	Groups []TenantGroupDetail `json:"groups,omitempty"`
}
