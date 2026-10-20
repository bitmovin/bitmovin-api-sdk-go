package model

// Complete snapshot of a Live Analysis at a specific result sequence
type AiSceneAnalysisLiveResult struct {
	// ID of the Live Analysis resource (required)
	AnalysisId *string `json:"analysisId,omitempty"`
	// ID of the Encoding associated with the Analysis (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Monotonically increasing result sequence, starting at 1 (required)
	Sequence *int64 `json:"sequence,omitempty"`
	// Time at which the AI analysis produced this result snapshot (required)
	ProducedAt *DateTime `json:"producedAt,omitempty"`
	// UTC presentation-clock time in Unix milliseconds corresponding to media time zero. This origin is fixed for the Live Analysis; coverage, observation, key-moment, and source-gap times remain relative to it. It is not the source capture time or result publication time. (required)
	MediaTimeOriginUnixMs *int64 `json:"mediaTimeOriginUnixMs,omitempty"`
	// Whether this is the final result snapshot. This does not by itself imply that the Analysis completed successfully. (required)
	IsFinal *bool `json:"isFinal,omitempty"`
	// Start of cumulative analyzed coverage on the monotonic analysis timeline (required)
	AnalyzedStartTimeSeconds *float64 `json:"analyzedStartTimeSeconds,omitempty"`
	// End of cumulative analyzed coverage on the monotonic analysis timeline (required)
	AnalyzedEndTimeSeconds *float64 `json:"analyzedEndTimeSeconds,omitempty"`
	// Cumulative closed source gaps on the monotonic analysis timeline. An empty or omitted array means there are no closed source gaps in this snapshot, not that the previous collection is unchanged.
	SourceGaps []AiSceneAnalysisLiveSourceGap `json:"sourceGaps,omitempty"`
	// Metadata for this result snapshot (required)
	Metadata *AiSceneAnalysisLiveResultMetadata `json:"metadata,omitempty"`
	// Complete set of active observations for this result sequence. An empty or omitted array means there are no active observations in this snapshot, not that the previous collection is unchanged. Observations may be removed in later sequences. Observation IDs identify immutable observation content and are never reused. Each time range identifies the analyzed media window that produced the observation, not an exact event location.
	Observations []AiSceneAnalysisLiveObservation `json:"observations,omitempty"`
	// Complete set of current key moments for this result sequence, ordered by ascending startInSeconds. Ordering of moments with equal start times is unspecified. An empty or omitted array means there are no current key moments in this snapshot, not that the previous collection is unchanged.  Moments may be refined or withdrawn in later result sequences until the final result. A refined moment retains its ID; a withdrawn moment is absent from subsequent snapshots and its ID is never reused. Consumers identify moments by ID and detect changes through lastUpdatedSequence, not array position. Changing array position alone or marking the result final does not change lastUpdatedSequence.  Every referenced observation is active in the same result snapshot. Moment references are updated or moments withdrawn when supporting observations are replaced or removed.
	KeyMoments []AiSceneAnalysisLiveKeyMoment `json:"keyMoments,omitempty"`
}
