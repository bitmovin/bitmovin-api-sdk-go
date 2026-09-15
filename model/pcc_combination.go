package model

// PccCombination model
type PccCombination struct {
	// The codec, as the shared contract spells it. (required)
	Codec *string `json:"codec,omitempty"`
	// The content protection, as the shared contract spells it. (required)
	Protection *string `json:"protection,omitempty"`
	// The column heading, spelled the way a reader reads it rather than the way the catalogue spells it. (required)
	Label *string `json:"label,omitempty"`
	// Whether this column's stream is HDR. Read it here rather than out of the codec name: not every HDR codec spells `hdr10`, and the Dolby Vision ones never do. (required)
	Hdr *bool `json:"hdr,omitempty"`
}
