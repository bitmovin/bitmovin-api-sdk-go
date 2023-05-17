package model

// LiveHlsManifest model
type LiveHlsManifest struct {
	// HLS manifest id (required)
	ManifestId *string `json:"manifestId,omitempty"`
	// Timeshift in seconds
	Timeshift *float64 `json:"timeshift,omitempty"`
	// Live edge offset in seconds
	LiveEdgeOffset *float64 `json:"liveEdgeOffset,omitempty"`
	// Specifies if the EXT-X-PROGRAM-DATETIME tag will be included
	InsertProgramDateTime *bool `json:"insertProgramDateTime,omitempty"`
	// Configuration for the EXT-X-PROGRAM-DATETIME tag
	ProgramDateTimeSettings *ProgramDateTimeSettings `json:"programDateTimeSettings,omitempty"`
}
