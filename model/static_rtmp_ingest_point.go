package model

// StaticRtmpIngestPoint model
type StaticRtmpIngestPoint struct {
	// The ID of the created static rtmp ingest point
	Id *string `json:"id,omitempty"`
	// Name of the ingest point. This can be helpful for easier identifying your ingest points
	Name                   *string                 `json:"name,omitempty"`
	StreamKeyConfiguration *StreamKeyConfiguration `json:"streamKeyConfiguration,omitempty"`
}
