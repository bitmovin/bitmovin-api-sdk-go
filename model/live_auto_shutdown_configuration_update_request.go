package model

// The auto shutdown configuration to set on a running Live Encoding. Any auto shutdown configuration the encoding currently has is overwritten.  Every timer that is omitted or `null` is disarmed. An empty object disarms all timers, so always send the complete configuration the encoding should run with, including the values that should stay unchanged.
type LiveAutoShutdownConfigurationUpdateRequest struct {
	// Automatically shutdown the live stream if there is no input anymore for a predefined number of seconds.
	BytesReadTimeoutSeconds *int64 `json:"bytesReadTimeoutSeconds,omitempty"`
	// Automatically stops the Live Encoding after the given number of minutes, counted from when the encoder applies this request. The maximum live encoding runtime of the organization, counted from the start of the encoding, still applies. If it comes first, the encoding stops at that limit instead.  The instant actually scheduled is returned as `scheduledShutdownAt`. Omit this field or set it to `null` to disarm this timer.
	StreamTimeoutMinutes *int64 `json:"streamTimeoutMinutes,omitempty"`
	// Automatically shutdown the live stream if input is never connected for a predefined number of minutes.
	WaitingForFirstConnectTimeoutMinutes *int64 `json:"waitingForFirstConnectTimeoutMinutes,omitempty"`
}
