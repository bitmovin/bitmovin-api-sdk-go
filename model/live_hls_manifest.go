package model

// LiveHlsManifest model
type LiveHlsManifest struct {
	// HLS manifest id (required)
	ManifestId *string `json:"manifestId,omitempty"`
	// Timeshift in seconds. We recommend to use a timeshift value not greater than 3 hours (10800.0 seconds). Longer values could negatively impact the manifest update frequency.
	Timeshift *float64 `json:"timeshift,omitempty"`
	// Live edge offset in seconds. New segments are held back for this duration (rounded up to whole segments) before they are added to the HLS playlist, and are always added in order. This adds the same amount of latency, but a segment that arrives late by up to this duration does not interrupt the playlist. A segment that arrives later than that stops the playlist from updating until it arrives. If it does not arrive at all, a placeholder segment is added in its place once a few newer segments are available. If set to 0, segments are added as soon as they are available. A late segment is then skipped and inserted into the playlist later, which can make players skip or repeat content. We recommend a value of at least 2 segment durations.
	LiveEdgeOffset *float64 `json:"liveEdgeOffset,omitempty"`
	// Specifies if the EXT-X-PROGRAM-DATETIME tag will be included
	InsertProgramDateTime *bool `json:"insertProgramDateTime,omitempty"`
	// Configuration for the EXT-X-PROGRAM-DATETIME tag
	ProgramDateTimeSettings *ProgramDateTimeSettings `json:"programDateTimeSettings,omitempty"`
	// Configuration for tags related to ad markers (e.g. Scte35)
	AdMarkerSettings *HlsManifestAdMarkerSettings `json:"adMarkerSettings,omitempty"`
}
