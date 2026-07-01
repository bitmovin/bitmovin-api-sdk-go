package model

// Manifest statistics for a live encoding heartbeat.
type LiveEncodingHeartbeatManifests struct {
	// Per-manifest HLS update statistics; one entry per HLS manifest.
	Hls []LiveEncodingHeartbeatHlsManifestStats `json:"hls,omitempty"`
}
