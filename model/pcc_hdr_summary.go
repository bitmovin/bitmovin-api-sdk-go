package model

// PccHdrSummary model
type PccHdrSummary struct {
	Outcomes []PccHdrOutcomeCount `json:"outcomes,omitempty"`
	// Selected device pools with at least one HDR playback pass. (required)
	ByDevice []PccHdrDevice `json:"byDevice,omitempty"`
	Shares   []PccHdrShare  `json:"shares,omitempty"`
	// HDR passes in the report. The picture question does not arise on a cell that failed. (required)
	Passes *float32 `json:"passes,omitempty"`
	// Passes carrying no reading at all, so a missing measurement never reads as a level of confidence. (required)
	Unreported *float32 `json:"unreported,omitempty"`
	// Passes where some instrument answered — frames read, or the device's own word. The `evidence` and `claim` shares above, added together. (required)
	Established *float32        `json:"established,omitempty"`
	Findings    []PccHdrFinding `json:"findings,omitempty"`
}
