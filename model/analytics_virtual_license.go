package model

// AnalyticsVirtualLicense model
type AnalyticsVirtualLicense struct {
	// Analytics Virtual License Key/Id
	Id *string `json:"id,omitempty"`
	// Name of the Analytics Virtual License
	Name *string `json:"name,omitempty"`
	// The timezone of the Analytics Virtual License
	Timezone *string `json:"timezone,omitempty"`
	// List of Analytics Licenses
	Licenses []AnalyticsVirtualLicenseLicensesListItem `json:"licenses,omitempty"`
	// The number of customData fields available
	CustomDataFieldsCount *int32 `json:"customDataFieldsCount,omitempty"`
	// Labels for Custom Data fields
	CustomDataFieldLabels *AnalyticsLicenseCustomDataFieldLabels `json:"customDataFieldLabels,omitempty"`
}
