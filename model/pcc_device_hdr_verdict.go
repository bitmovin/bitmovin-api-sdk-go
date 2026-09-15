package model

// PccDeviceHdrVerdict model
type PccDeviceHdrVerdict struct {
	Mark       *string          `json:"mark,omitempty"`
	Label      *string          `json:"label,omitempty"`
	Confidence PccHdrConfidence `json:"confidence,omitempty"`
	Actionable *bool            `json:"actionable,omitempty"`
	// HDR passes supporting this verdict, not all passes for the device. (required)
	Passes *float32 `json:"passes,omitempty"`
}
