package model

// AnalyticsVirtualLicenseRequest model
type AnalyticsVirtualLicenseRequest struct {
	// Name of the Analytics Virtual License (required)
	Name *string `json:"name,omitempty"`
	// The timezone of the Analytics Virtual License (required)
	Timezone *string `json:"timezone,omitempty"`
	// List of Analytics Licenses (required)
	Licenses []AnalyticsVirtualLicenseLicensesListItem `json:"licenses,omitempty"`
}
