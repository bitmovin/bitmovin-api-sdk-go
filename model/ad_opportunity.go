package model

// AdOpportunity model
type AdOpportunity struct {
	// The reason why the scene was rated with a certain score
	Reason *string `json:"reason,omitempty"`
	// Score from 0.0 to 1.0 rating the ad placement suitability at the end of a scene based on content analysis
	Score *float32 `json:"score,omitempty"`
}
