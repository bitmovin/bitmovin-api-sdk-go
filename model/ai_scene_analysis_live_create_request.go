package model

// Configuration for a Live Analysis. Each recording and analysis destination references an existing Encoding Output or provides an inline Output definition. Inline Outputs are created synchronously. Within this request, identical complete inline Output definitions, including credentials, are created once and reused across all destinations; destination paths and ACLs do not affect that reuse.
type AiSceneAnalysisLiveCreateRequest struct {
	// Name of the Analysis
	Name *string `json:"name,omitempty"`
	// Key used to publish the RTMP stream. When the Live Analysis is `RUNNING`, the Get Live Analysis details response returns the current value in `ingest.streamKey`. (required)
	StreamKey *string `json:"streamKey,omitempty"`
	// Region in which the AI analysis runs. `EXTERNAL` is not supported yet.
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Destinations for the stream recording (required)
	Recording *AiSceneAnalysisLiveRecordingRequest `json:"recording,omitempty"`
	// Destinations for cumulative AI analysis results (required)
	Outputs []AiSceneAnalysisLiveOutput `json:"outputs,omitempty"`
}
