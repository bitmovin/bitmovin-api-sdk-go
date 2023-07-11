package model

// StreamsStyleConfigResponse model
type StreamsStyleConfigResponse struct {
	// The identifier of the style config
	Id *string `json:"id,omitempty"`
	// UUID of the associated organization
	OrgId       *string                        `json:"orgId,omitempty"`
	PlayerStyle *StreamsStyleConfigPlayerStyle `json:"playerStyle,omitempty"`
	// URL of the watermark image
	WatermarkUrl *string `json:"watermarkUrl,omitempty"`
	// Target link of the watermark image
	WatermarkTargetLink *string `json:"watermarkTargetLink,omitempty"`
}
