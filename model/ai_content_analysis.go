package model

// AiContentAnalysis model
type AiContentAnalysis struct {
	// AI service settings
	AiService *AiService `json:"aiService,omitempty"`
	// Features of the AI content analysis
	Features *AiContentAnalysisFeatures `json:"features,omitempty"`
}
