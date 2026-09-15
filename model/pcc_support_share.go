package model

// PccSupportShare model
type PccSupportShare struct {
	// What this share counts, in words rather than in codes. (required)
	Label *string `json:"label,omitempty"`
	// Device pools that played it. (required)
	Played *float32 `json:"played,omitempty"`
	// Device pools that answered either way. The denominator, never the fleet size. (required)
	Measured *float32 `json:"measured,omitempty"`
}
