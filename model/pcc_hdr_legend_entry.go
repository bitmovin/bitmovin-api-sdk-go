package model

// PccHdrLegendEntry model
type PccHdrLegendEntry struct {
	// The mark an HDR result wears on the cell. (required)
	Mark *string `json:"mark,omitempty"`
	// The reader's word for it — `The frames really were HDR`, `The device claims HDR`, and so on. (required)
	Label *string `json:"label,omitempty"`
	// Which instrument established the picture, where anything did. (required)
	Confidence PccHdrConfidence `json:"confidence,omitempty"`
	// What that mark establishes, and what it does not. (required)
	Sentence *string `json:"sentence,omitempty"`
	// Whether a result wearing this mark is one somebody should chase, which is not the same as how confident it is. (required)
	Actionable *bool `json:"actionable,omitempty"`
}
