package model

// PccHdrDevice model
type PccHdrDevice struct {
	// Opaque stable pool key for row identity. Do not display it as a device name. (required)
	Key       *string `json:"key,omitempty"`
	Name      *string `json:"name,omitempty"`
	Qualifier *string `json:"qualifier,omitempty"`
	// All selected HDR playback passes for this device pool. (required)
	Passes *float32 `json:"passes,omitempty"`
	// Service-resolved verdict: negative findings outrank positive results. Null when no pass carries an HDR reading. (required)
	Verdict  *PccDeviceHdrVerdict `json:"verdict,omitempty"`
	Outcomes []PccHdrOutcomeCount `json:"outcomes,omitempty"`
}
