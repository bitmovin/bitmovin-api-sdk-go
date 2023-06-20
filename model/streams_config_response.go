package model

// StreamsConfigResponse model
type StreamsConfigResponse struct {
	// The identifier of the stream config
	Id *string `json:"id,omitempty"`
	// UUID of the associated organization
	OrgId       *string                   `json:"orgId,omitempty"`
	PlayerStyle *StreamsConfigPlayerStyle `json:"playerStyle,omitempty"`
	// URL of the watermark image
	WatermarkUrl *string `json:"watermarkUrl,omitempty"`
	// Target link of the watermark image
	WatermarkTargetLink *string `json:"watermarkTargetLink,omitempty"`
}
