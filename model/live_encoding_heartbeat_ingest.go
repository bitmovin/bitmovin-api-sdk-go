package model

// Information about the live ingestion status
type LiveEncodingHeartbeatIngest struct {
	Status LiveEncodingStatus `json:"status,omitempty"`
	// Indicates whether the ingest is healthy.
	Healthy *bool `json:"healthy,omitempty"`
	// Data about individual ingestPoints within the active live ingest.
	IngestPoints []LiveEncodingHeartbeatIngestPoint `json:"ingestPoints,omitempty"`
	// Data about individual streams within the active live ingest.
	Streams            []LiveEncodingHeartbeatIngestStream `json:"streams,omitempty"`
	RtmpUserIngestInfo *RtmpUserIngestInfo                 `json:"rtmpUserIngestInfo,omitempty"`
}
