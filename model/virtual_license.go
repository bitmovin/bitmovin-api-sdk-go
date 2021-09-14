package model

// VirtualLicense model
type VirtualLicense struct {
	// Virtual License Key/Id (required)
	Id *string `json:"id,omitempty"`
	// Name of the Virtual License
	Name *string `json:"name,omitempty"`
	// The timezone of the Virtual License
	Timezone *string `json:"timezone,omitempty"`
	// List of Analytics Licenses
	Licenses []VirtualLicenseLicensesListItem `json:"licenses,omitempty"`
}
