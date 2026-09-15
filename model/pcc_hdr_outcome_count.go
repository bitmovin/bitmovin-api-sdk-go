package model

// PccHdrOutcomeCount model
type PccHdrOutcomeCount struct {
	Outcome PccHdrOutcome `json:"outcome,omitempty"`
	// Null only for unreported readings. Evidence includes both HDR and SDR outcomes. (required)
	Confidence PccHdrConfidence `json:"confidence,omitempty"`
	Label      *string          `json:"label,omitempty"`
	// Selected HDR playback passes with this outcome. All six outcomes sum to HDR passes. (required)
	Cells *float32 `json:"cells,omitempty"`
}
