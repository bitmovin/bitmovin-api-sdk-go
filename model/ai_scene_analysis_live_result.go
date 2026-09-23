package model

// Cumulative immutable result generation for a Live Analysis
type AiSceneAnalysisLiveResult struct {
	// ID of the Live Analysis resource (required)
	AnalysisId *string `json:"analysisId,omitempty"`
	// ID of the Encoding associated with the Analysis (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Monotonically increasing generation sequence, starting at 1 (required)
	Sequence *int64 `json:"sequence,omitempty"`
	// Time at which the AI analysis produced this result generation (required)
	ProducedAt *DateTime `json:"producedAt,omitempty"`
	// Whether AI analysis produced this as the final result generation. This does not by itself imply that the Analysis completed successfully. (required)
	IsFinal *bool `json:"isFinal,omitempty"`
	// Start of cumulative analyzed coverage on the monotonic analysis timeline (required)
	AnalyzedStartTimeSeconds *float64 `json:"analyzedStartTimeSeconds,omitempty"`
	// End of cumulative analyzed coverage on the monotonic analysis timeline (required)
	AnalyzedEndTimeSeconds *float64 `json:"analyzedEndTimeSeconds,omitempty"`
	// Cumulative closed source gaps on the monotonic analysis timeline (required)
	SourceGaps []AiSceneAnalysisLiveSourceGap `json:"sourceGaps,omitempty"`
	// Producer metadata for this generation (required)
	Metadata *AiSceneAnalysisLiveResultMetadata `json:"metadata,omitempty"`
	// Cumulative immutable observations. Existing observations retain the same ID and content across later generations. Each time range identifies the analyzed media window that produced the observation, not an exact event location. (required)
	Observations []AiSceneAnalysisLiveObservation `json:"observations,omitempty"`
}
