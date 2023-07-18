package model

// LiveEncodingOptionsStatistics model
type LiveEncodingOptionsStatistics struct {
	// The ID of the encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Live option units used (required)
	UnitsUsed *float64        `json:"unitsUsed,omitempty"`
	Type      LiveOptionsType `json:"type,omitempty"`
}
