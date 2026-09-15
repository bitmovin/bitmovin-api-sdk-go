package model

// PccVerdictLegendEntry model
type PccVerdictLegendEntry struct {
	// The mark the grid draws for every verdict reading as this entry's label. (required)
	Symbol *string `json:"symbol,omitempty"`
	// The reader's word for those verdicts — `Supported`, `Not measured`, and so on. (required)
	Label *string `json:"label,omitempty"`
	// What that word means here. (required)
	Sentence *string `json:"sentence,omitempty"`
	// False where these verdicts say something about the measurement rather than about the device. Folding those into \"not supported\" is how this dataset gets misread. (required)
	AboutTheDevice *bool `json:"aboutTheDevice,omitempty"`
}
