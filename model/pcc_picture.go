package model

// PccPicture model
type PccPicture struct {
	// The mark the cell wears. `hdrLegend` explains it. (required)
	Mark       *string          `json:"mark,omitempty"`
	Confidence PccHdrConfidence `json:"confidence,omitempty"`
	// Whether this is a result somebody should chase, which is not how confident it is. (required)
	Actionable *bool   `json:"actionable,omitempty"`
	Sentence   *string `json:"sentence,omitempty"`
}
