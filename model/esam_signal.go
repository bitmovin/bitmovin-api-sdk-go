package model

// ESAM signal following the SCTE-250 standard
type EsamSignal struct {
	// The offset from the matched signal in ISO 8601 duration format, accurate to milliseconds
	Offset *string `json:"offset,omitempty"`
	// Base64-encoded SCTE-35 binary data to be inserted into the stream (required)
	Binary *string `json:"binary,omitempty"`
	// Interval in ISO 8601 duration format for which the signal should be repeated.  A signal may be specified as repeating when the interval and end attributes are present.  In this case, the signal is executed at the time specified by offset and again at the time  specified by adding interval to offset. This should be continued until reaching the duration  of offset + end.
	Interval *string `json:"interval,omitempty"`
	// End duration in ISO 8601 duration format when a repeated signal should stop being repeated
	End *string `json:"end,omitempty"`
}
