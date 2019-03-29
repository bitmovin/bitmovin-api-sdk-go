package model

type AudioGroupConfiguration struct {
	// Dropping mode
	DroppingMode VariantStreamDroppingMode `json:"droppingMode,omitempty"`
	// Audio groups
	Groups []AudioGroup `json:"groups,omitempty"`
}

