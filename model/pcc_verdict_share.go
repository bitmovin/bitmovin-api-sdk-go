package model

// PccVerdictShare model
type PccVerdictShare struct {
	Verdict PccVerdict `json:"verdict,omitempty"`
	// The reader's word for it. (required)
	Label *string `json:"label,omitempty"`
	// Cells sharing this label, including verdicts grouped under `Not measured`. (required)
	Cells *float32 `json:"cells,omitempty"`
	// False where the verdict says something about the measurement, not about the device. (required)
	AboutTheDevice *bool `json:"aboutTheDevice,omitempty"`
}
