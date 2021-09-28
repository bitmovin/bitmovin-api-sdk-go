package model

// AnalyticsVirtualLicenseLicensesListItem model
type AnalyticsVirtualLicenseLicensesListItem struct {
	// Analytics License Id
	Id *string `json:"id,omitempty"`
	// Analytics License key
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Organisation Id of license
	OrgId *string `json:"orgId,omitempty"`
	// Analytics License name
	Name *string `json:"name,omitempty"`
}
