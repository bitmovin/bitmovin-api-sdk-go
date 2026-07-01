package model

// A single per-muxing segment upload whose duration exceeded the upload-outlier threshold.
type LiveEncodingHeartbeatUploadOutlier struct {
	// Output segment number this upload belongs to.
	SegmentNumber *int64 `json:"segmentNumber,omitempty"`
	// Identifier of the muxing whose upload was flagged.
	EncodingReferenceMuxingId *string `json:"encodingReferenceMuxingId,omitempty"`
	// Upload duration in milliseconds.
	DurationMs *int64 `json:"durationMs,omitempty"`
}
