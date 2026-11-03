package model

// AdPosition model
type AdPosition struct {
	// Selected ad break position in seconds from the beginning of the video.
	Position *float64 `json:"position,omitempty"`
	// Ad duration in seconds requested for this placement. Not set for breaks added at all scene boundaries, when the schedule entry has no duration, or when the encoder does not support ad durations.
	Duration *float64 `json:"duration,omitempty"`
}
