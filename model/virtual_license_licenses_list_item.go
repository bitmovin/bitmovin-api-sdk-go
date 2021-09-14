package model

// VirtualLicenseLicensesListItem model
type VirtualLicenseLicensesListItem struct {
	// Analytics License Id
	Id *string `json:"id,omitempty"`
	// Analytics License key
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Organisation Id of license
	OrgId *string `json:"orgId,omitempty"`
}
