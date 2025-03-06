package model

// AutomaticAdPlacementPosition model
type AutomaticAdPlacementPosition struct {
	// Position of the ad placement in seconds.
	Position *float64 `json:"position,omitempty"`
	// Maximum deviation in seconds to the ad placement position.
	MaxDeviation *float64 `json:"maxDeviation,omitempty"`
	// The desired duration in seconds of the ad to be inserted.
	Duration *float64 `json:"duration,omitempty"`
}
