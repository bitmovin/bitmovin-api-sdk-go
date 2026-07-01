package model

// A single media-advancing HLS manifest update.
type LiveEncodingHeartbeatManifestUpdate struct {
	// Wall-clock time the media-advancing manifest update occured, returned as UTC in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ManifestUpdateWallClockTime *DateTime `json:"manifestUpdateWallClockTime,omitempty"`
	// Latest media presentation time across renditions (min of all playlists' stream progress), returned as UTC in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	LatestMediaPresentationTime *DateTime `json:"latestMediaPresentationTime,omitempty"`
	// Manifest update latency in milliseconds (elapsed wall-clock minus media-time advanced).
	ManifestUpdateLatency *int64 `json:"manifestUpdateLatency,omitempty"`
}
