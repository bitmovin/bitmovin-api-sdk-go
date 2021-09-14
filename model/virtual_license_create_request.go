package model

// VirtualLicenseCreateRequest model
type VirtualLicenseCreateRequest struct {
	// Virtual License Key/Id (required)
	Id *string `json:"id,omitempty"`
	// Name of the Virtual License (required)
	Name *string `json:"name,omitempty"`
	// The timezone of the Virtual License (required)
	Timezone *string `json:"timezone,omitempty"`
	// List of Analytics Licenses
	Licenses []VirtualLicenseLicensesListItem `json:"licenses,omitempty"`
}
