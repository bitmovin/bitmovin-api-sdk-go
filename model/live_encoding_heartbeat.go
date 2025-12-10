package model

// Heartbeat data for a Live Encoding.
type LiveEncodingHeartbeat struct {
	// timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	Timestamp *DateTime `json:"timestamp,omitempty"`
	// Information about the live ingestion status
	Ingest *LiveEncodingHeartbeatIngest `json:"ingest,omitempty"`
	// Live encoding heartbeat events
	Events []LiveEncodingHeartbeatEvent `json:"events,omitempty"`
}
