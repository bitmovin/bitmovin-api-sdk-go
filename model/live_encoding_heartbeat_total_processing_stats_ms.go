package model

// Aggregate statistics of per-segment total processing duration over a rolling window of completed segments.
type LiveEncodingHeartbeatTotalProcessingStatsMs struct {
	// Minimum total processing duration in milliseconds.
	Min *int64 `json:"min,omitempty"`
	// Maximum total processing duration in milliseconds.
	Max *int64 `json:"max,omitempty"`
	// Mean total processing duration in milliseconds.
	Mean *int64 `json:"mean,omitempty"`
	// Median total processing duration in milliseconds.
	Median *int64 `json:"median,omitempty"`
}
