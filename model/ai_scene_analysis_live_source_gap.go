package model

// AiSceneAnalysisLiveSourceGap model
type AiSceneAnalysisLiveSourceGap struct {
	// Gap start on the monotonic analysis timeline (required)
	StartTimeSeconds *float64 `json:"startTimeSeconds,omitempty"`
	// Gap end on the monotonic analysis timeline (required)
	EndTimeSeconds *float64 `json:"endTimeSeconds,omitempty"`
	// Reason for the source gap (required)
	Reason AiSceneAnalysisLiveSourceGapReason `json:"reason,omitempty"`
}
