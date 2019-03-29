package model

type RtmpIngestPoint struct {
	// The name of the application where the ingest is streamed to. This has to be unique for each ingest point
	ApplicationName string `json:"applicationName,omitempty"`
	// The stream key for the backup input
	StreamKey string `json:"streamKey,omitempty"`
}

