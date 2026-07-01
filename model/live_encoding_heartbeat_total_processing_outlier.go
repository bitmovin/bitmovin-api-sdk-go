package model

// A single segment whose total processing duration exceeded twice the rolling median.
type LiveEncodingHeartbeatTotalProcessingOutlier struct {
	// Output segment number this measurement belongs to.
	SegmentNumber *int64 `json:"segmentNumber,omitempty"`
	// Total processing duration in milliseconds.
	DurationMs *int64 `json:"durationMs,omitempty"`
}
