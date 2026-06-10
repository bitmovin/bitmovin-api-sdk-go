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
	// Total number of dropped video packets since the live encoding started.
	DroppedPacketsVideo *int64 `json:"droppedPacketsVideo,omitempty"`
	// Total number of dropped audio packets since the live encoding started.
	DroppedPacketsAudio *int64 `json:"droppedPacketsAudio,omitempty"`
	// Total number of corrupt video packets since the live encoding started.
	CorruptPacketsVideo *int64 `json:"corruptPacketsVideo,omitempty"`
	// Total number of corrupt audio packets since the live encoding started.
	CorruptPacketsAudio *int64 `json:"corruptPacketsAudio,omitempty"`
}
