package model

// Period model
type Period struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Starting time in seconds
	Start *float64 `json:"start,omitempty"`
	// Duration in seconds.<br/>Please note that the duration of a Period is usually determined by the media contained therein.<br/>Setting the `duration` property to a specific value will override this default behaviour.<br/>Warning: Use at your own risk!
	Duration *float64 `json:"duration,omitempty"`
}
