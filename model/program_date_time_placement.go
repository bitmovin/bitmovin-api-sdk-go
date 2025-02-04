package model

// ProgramDateTimePlacement model
type ProgramDateTimePlacement struct {
	ProgramDateTimePlacementMode ProgramDateTimePlacementMode `json:"programDateTimePlacementMode,omitempty"`
	// Interval for placing ProgramDateTime. Only required for SEGMENTS_INTERVAL or SECONDS_INTERVAL.
	Interval *int32 `json:"interval,omitempty"`
}
