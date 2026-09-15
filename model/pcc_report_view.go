package model

// PccReportView model
type PccReportView struct {
	// Effective trimmed lowercase substring matched against the published device name and qualifier. Empty means all devices. (required)
	Device *string `json:"device,omitempty"`
	// Effective trimmed lowercase substring matched against codec identifiers. Empty means all codecs. (required)
	Codec *string `json:"codec,omitempty"`
	// Whether only HDR columns are selected. (required)
	HdrOnly *bool `json:"hdrOnly,omitempty"`
	// Whether pools need at least one selected cell answering about the device. (required)
	ReportedOnly *bool `json:"reportedOnly,omitempty"`
	// Whether pools with only prerelease browser evidence are included. (required)
	IncludePrerelease *bool `json:"includePrerelease,omitempty"`
}
