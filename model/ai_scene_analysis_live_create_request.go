package model

// Configuration for a Live Analysis. Each recording and analysis destination references an existing Encoding Output or provides an inline Output definition. Inline Outputs are created synchronously. Within this request, identical complete inline Output definitions, including credentials, are created once and reused across all destinations; destination paths and ACLs do not affect that reuse.
type AiSceneAnalysisLiveCreateRequest struct {
	// Name of the Analysis
	Name *string `json:"name,omitempty"`
	// Key used to publish the RTMP stream. When the Live Analysis is `RUNNING`, the Get Live Analysis details response returns the current value in `ingest.streamKey`. (required)
	StreamKey *string `json:"streamKey,omitempty"`
	// Region in which the AI analysis runs. `EXTERNAL` is not supported yet.
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Encoder release for the backing Live Encoding, using the same values and validation as `Encoding.encoderVersion` (for example, `LATEST` or a supported release version). Passed unchanged to the Encoding Resource Service. If omitted, the server default is used. Selects the full encoder release; component header overrides remain separate.
	EncoderVersion *string `json:"encoderVersion,omitempty"`
	// Destinations for the stream recording (required)
	Recording *AiSceneAnalysisLiveRecordingRequest `json:"recording,omitempty"`
	// Destinations for incremental Live Analysis result batches (required)
	Outputs []AiSceneAnalysisLiveOutput `json:"outputs,omitempty"`
}
