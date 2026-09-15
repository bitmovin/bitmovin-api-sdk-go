package model

// PccHdrShare model
type PccHdrShare struct {
	Confidence PccHdrConfidence `json:"confidence,omitempty"`
	Label      *string          `json:"label,omitempty"`
	Cells      *float32         `json:"cells,omitempty"`
}
