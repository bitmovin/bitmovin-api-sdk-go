package model

// UpdateEncodingRtmpIngestPointRequest model
type UpdateEncodingRtmpIngestPointRequest struct {
	// List of ingest points to be updated for the encoding. The RTMPs application name and streamKey will be assigned to the encoding.
	IngestPoints []RtmpIngestPoint `json:"ingestPoints,omitempty"`
}
