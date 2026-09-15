package model

// PccCodecReach model
type PccCodecReach struct {
	Codec        *string                   `json:"codec,omitempty"`
	ByDeviceType []PccCodecDeviceTypeReach `json:"byDeviceType,omitempty"`
}
