package model

// PccCodecDeviceTypeReach model
type PccCodecDeviceTypeReach struct {
	DeviceType *string `json:"deviceType,omitempty"`
	// Distinct selected device pools that played this codec. (required)
	Played *float32 `json:"played,omitempty"`
	// Distinct selected pools that answered about this codec. A protection-only refusal is excluded. (required)
	Measured *float32 `json:"measured,omitempty"`
}
