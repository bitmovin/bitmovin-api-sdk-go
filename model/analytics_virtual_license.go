package model

// AnalyticsVirtualLicense model
type AnalyticsVirtualLicense struct {
	// Analytics Virtual License Key/Id
	Id *string `json:"id,omitempty"`
	// Name of the Analytics Virtual License
	Name *string `json:"name,omitempty"`
	// The timezone of the Analytics Virtual License
	Timezone *string `json:"timezone,omitempty"`
	// Retention time of impressions, returned as ISO 8601 duration format: P(n)Y(n)M(n)DT(n)H(n)M(n)S
	RetentionTime *string `json:"retentionTime,omitempty"`
	// List of Analytics Licenses
	Licenses []AnalyticsVirtualLicenseLicensesListItem `json:"licenses,omitempty"`
	// The number of customData fields available
	CustomDataFieldsCount *int32 `json:"customDataFieldsCount,omitempty"`
	// Labels for Custom Data fields
	CustomDataFieldLabels *AnalyticsLicenseCustomDataFieldLabels `json:"customDataFieldLabels,omitempty"`
	// The expiration date of the license if applicable, returned as ISO 8601 date-time format
	PlanExpiredAt *DateTime `json:"planExpiredAt,omitempty"`
}
