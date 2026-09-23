package model

// Immutable consumer-visible observation produced from an analyzed media window
type AiSceneAnalysisLiveObservation struct {
	// Stable opaque observation ID that remains unchanged across cumulative result generations (required)
	Id *string `json:"id,omitempty"`
	// Consumer-visible description of a development in the analyzed media (required)
	Text *string `json:"text,omitempty"`
	// Start of the analyzed media window that produced the observation (required)
	StartTimeSeconds *float64 `json:"startTimeSeconds,omitempty"`
	// End of the analyzed media window that produced the observation (required)
	EndTimeSeconds *float64 `json:"endTimeSeconds,omitempty"`
}
