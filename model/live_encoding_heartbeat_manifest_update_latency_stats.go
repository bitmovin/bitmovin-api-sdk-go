package model

// Aggregate statistics of HLS manifest update latencies.
type LiveEncodingHeartbeatManifestUpdateLatencyStats struct {
	// Mean manifest update latency in milliseconds.
	Mean *int64 `json:"mean,omitempty"`
	// Median manifest update latency in milliseconds.
	Median *int64 `json:"median,omitempty"`
	// Minimum manifest update latency in milliseconds.
	Min *int64 `json:"min,omitempty"`
	// Maximum manifest update latency in milliseconds.
	Max *int64 `json:"max,omitempty"`
}
