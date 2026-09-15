package model

// PccReport model
type PccReport struct {
	// Identifies this report. A new generation produces a new one. (required)
	ReportId *string `json:"reportId,omitempty"`
	// When the generation that produced this report finished. (required)
	GeneratedAt *string `json:"generatedAt,omitempty"`
	// The selected report. By default, pools whose recorded browsers are all pre-release are excluded; includePrerelease retains them. Unknown browser identities do not cause prerelease exclusion. Null where a report is held that this service cannot read, which a generation replaces. (required)
	Report *PccCompatibilityMatrix `json:"report,omitempty"`
}
