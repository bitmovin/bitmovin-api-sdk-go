package model

// PccDeviceTypeShare model
type PccDeviceTypeShare struct {
	// The device type reported by the fleet. (required)
	Label *string `json:"label,omitempty"`
	// Cells that played, across device pools of this kind. (required)
	Played *float32 `json:"played,omitempty"`
	// Cells with a device-answering verdict, across pools of this kind. The denominator. (required)
	Measured *float32 `json:"measured,omitempty"`
	// Device pools of this kind with at least one device-answering verdict. (required)
	Models *float32 `json:"models,omitempty"`
}
