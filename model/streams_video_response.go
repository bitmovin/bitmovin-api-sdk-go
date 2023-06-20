package model

// StreamsVideoResponse model
type StreamsVideoResponse struct {
	// The identifier of the stream
	Id *string `json:"id,omitempty"`
	// The asset URL of the stream
	AssetUrl *string `json:"assetUrl,omitempty"`
	// The title of the stream
	Title *string `json:"title,omitempty"`
	// The description of the stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// The status of the stream
	Status StreamsVideoStatus     `json:"status,omitempty"`
	Config *StreamsConfigResponse `json:"config,omitempty"`
	// List of encoding status information
	EncodingTasks []StreamsVideoEncodingTask `json:"encodingTasks,omitempty"`
	// Poster URL
	PosterUrl         *string                           `json:"posterUrl,omitempty"`
	AdConfig          *StreamsAdConfigResponse          `json:"adConfig,omitempty"`
	ContentProtection *StreamsContentProtectionResponse `json:"contentProtection,omitempty"`
}
