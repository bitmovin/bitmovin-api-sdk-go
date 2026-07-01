package model

// HLS manifest update monitoring data for a live encoding heartbeat.
type LiveEncodingHeartbeatHlsManifestStats struct {
	// Id of the HLS manifest these statistics belong to.
	ManifestId *string `json:"manifestId,omitempty"`
	// Aggregate latency statistics over recorded manifest updates.
	ManifestUpdateLatencyStats *LiveEncodingHeartbeatManifestUpdateLatencyStats `json:"manifestUpdateLatencyStats,omitempty"`
	// Per-manifest outlier-only (high-latency) manifest updates, ordered newest to oldest.
	ManifestUpdates []LiveEncodingHeartbeatManifestUpdate `json:"manifestUpdates,omitempty"`
}
