package model

// LiveAutoShutdownConfiguration model
type LiveAutoShutdownConfiguration struct {
	// Automatically shutdown the live stream if there is no input anymore for a predefined number of seconds.
	BytesReadTimeoutSeconds *int64 `json:"bytesReadTimeoutSeconds,omitempty"`
	// Automatically shutdown the live stream after a predefined runtime in minutes.
	StreamTimeoutMinutes *int64 `json:"streamTimeoutMinutes,omitempty"`
}
