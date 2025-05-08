package model

// AiSceneAnalysis model
type AiSceneAnalysis struct {
	// AI service settings
	AiService *AiService `json:"aiService,omitempty"`
	// Features of the AI scene analysis
	Features *AiSceneAnalysisFeatures `json:"features,omitempty"`
}
