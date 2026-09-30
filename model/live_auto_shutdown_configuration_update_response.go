package model

// The auto shutdown configuration that was accepted and applied to the running Live Encoding, together with the instant the encoder armed the shutdown for.  `streamTimeoutMinutes` was applied from the moment the encoder accepted the update, not from the start of the encoding, so `scheduledShutdownAt` rather than the timeout value is what says when this encoding stops.
type LiveAutoShutdownConfigurationUpdateResponse struct {
	// Automatically shutdown the live stream if there is no input anymore for a predefined number of seconds.
	BytesReadTimeoutSeconds *int64 `json:"bytesReadTimeoutSeconds,omitempty"`
	// Automatically shutdown the live stream after a predefined runtime in minutes.
	StreamTimeoutMinutes *int64 `json:"streamTimeoutMinutes,omitempty"`
	// Automatically shutdown the live stream if input is never connected for a predefined number of minutes.
	WaitingForFirstConnectTimeoutMinutes *int64 `json:"waitingForFirstConnectTimeoutMinutes,omitempty"`
	// The instant at which the Live Encoding is currently scheduled to shut down, as reported by the encoder. `null` means no shutdown is scheduled.  `bytesReadTimeoutSeconds` is not reflected here, as it only arms once the input stops flowing, so the encoding can still shut down earlier than this.
	ScheduledShutdownAt *DateTime `json:"scheduledShutdownAt,omitempty"`
}
