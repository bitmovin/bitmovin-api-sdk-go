package model

// AiSceneAnalysisRegulatoryAdvisories model
type AiSceneAnalysisRegulatoryAdvisories struct {
	// The regulatory advisory topics to screen the asset for. At least one topic must be set. (required)
	Topics []RegulatoryAdvisoryTopic `json:"topics,omitempty"`
}
