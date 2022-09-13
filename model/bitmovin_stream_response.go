package model

// BitmovinStreamResponse model
type BitmovinStreamResponse struct {
	// The identifier of the Stream
	Id *string `json:"id,omitempty"`
	// The asset URL of the Stream
	AssetUrl *string `json:"assetUrl,omitempty"`
	// The title of the Stream
	Title *string `json:"title,omitempty"`
	// The description of the Stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// The status of the Stream
	Status BitmovinStreamStatus `json:"status,omitempty"`
	// The target quality of the Stream
	TargetQuality BitmovinStreamQuality `json:"targetQuality,omitempty"`
	// List of available stream qualities
	AvailableQualities []BitmovinStreamQuality `json:"availableQualities,omitempty"`
}
