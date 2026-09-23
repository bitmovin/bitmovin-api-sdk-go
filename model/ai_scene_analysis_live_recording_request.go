package model

// AiSceneAnalysisLiveRecordingRequest model
type AiSceneAnalysisLiveRecordingRequest struct {
	// Destinations for the stream recording (required)
	Outputs []AiSceneAnalysisLiveOutput `json:"outputs,omitempty"`
}
