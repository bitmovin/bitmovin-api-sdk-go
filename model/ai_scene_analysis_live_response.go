package model

// AiSceneAnalysisLiveResponse model
type AiSceneAnalysisLiveResponse struct {
	// ID of the Live Analysis resource (required)
	AnalysisId *string `json:"analysisId,omitempty"`
	// ID of the Encoding associated with the Analysis (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Name of the Analysis
	Name *string `json:"name,omitempty"`
	// Current lifecycle state of the Live Analysis (required)
	Status AiSceneAnalysisLiveStatus `json:"status,omitempty"`
	// Resolved output configuration for the stream recording (required)
	Recording *AiSceneAnalysisLiveRecording `json:"recording,omitempty"`
	// Resolved Encoding Output ID references for cumulative AI analysis results (required)
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Current RTMP ingest details. Present only in the Get Live Analysis details response while the Live Analysis is `RUNNING`.
	Ingest *LiveEncoding `json:"ingest,omitempty"`
	// Failure details. Present only when the status is `ERROR` or `TRANSFER_ERROR`.
	Error *AiSceneAnalysisLiveError `json:"error,omitempty"`
	// Creation timestamp, returned as UTC in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	CreatedAt *DateTime `json:"createdAt,omitempty"`
}
