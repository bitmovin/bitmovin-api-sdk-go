package model

// Segment processing statistics for a live encoding heartbeat.
type LiveEncodingHeartbeatSegments struct {
	// Aggregate statistics of per-segment total processing duration.
	TotalProcessingStatsMs *LiveEncodingHeartbeatTotalProcessingStatsMs `json:"totalProcessingStatsMs,omitempty"`
	// Last 20 per-muxing uploads whose duration exceeded the upload-outlier threshold.
	UploadOutliers []LiveEncodingHeartbeatUploadOutlier `json:"uploadOutliers,omitempty"`
	// Last 20 segments whose total processing duration exceeded twice the rolling median.
	TotalProcessingOutliers []LiveEncodingHeartbeatTotalProcessingOutlier `json:"totalProcessingOutliers,omitempty"`
}
