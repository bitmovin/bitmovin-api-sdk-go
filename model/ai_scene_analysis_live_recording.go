package model

// Resolved output configuration for the stream recording
type AiSceneAnalysisLiveRecording struct {
	// Resolved Encoding Output ID references for the stream recording (required)
	Outputs []EncodingOutput `json:"outputs,omitempty"`
}
