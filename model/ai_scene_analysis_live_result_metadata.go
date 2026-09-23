package model

// AiSceneAnalysisLiveResultMetadata model
type AiSceneAnalysisLiveResultMetadata struct {
	// Version of the AI analysis software (required)
	Version *string `json:"version,omitempty"`
	// Disclaimer associated with AI-generated analysis data (required)
	Disclaimer *string `json:"disclaimer,omitempty"`
}
