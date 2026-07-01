package model

// Output statistics for a live encoding heartbeat.
type LiveEncodingHeartbeatOutput struct {
	// Manifest statistics for the live encoding output.
	Manifests *LiveEncodingHeartbeatManifests `json:"manifests,omitempty"`
	// Segment processing statistics for the live encoding output.
	Segments *LiveEncodingHeartbeatSegments `json:"segments,omitempty"`
}
