package model

// ESAM media point following the SCTE-250 standard
type EsamMediaPoint struct {
	// ISO 8601 date-time specifying when the signal should be matched and inserted into the live stream
	MatchTime *DateTime `json:"matchTime,omitempty"`
	// Array of ESAM signals containing SCTE-35 binary data. At least one signal is required (required)
	Signals []EsamSignal `json:"signals,omitempty"`
	// Optional array of ESAM conditions with timing offsets and directional markers (OUT/IN) for signaling content boundaries
	Conditions []EsamCondition `json:"conditions,omitempty"`
}
