package model

// AdOpportunity model
type AdOpportunity struct {
	// Explanation of the scene boundary's suitability for an ad break, including low or zero suitability.
	Reason *string `json:"reason,omitempty"`
	// Suitability of placing an ad break at the end of the scene. Higher scores indicate greater suitability. A score of 0.5 can also mean that no assessment is available.
	Score *float64 `json:"score,omitempty"`
}
