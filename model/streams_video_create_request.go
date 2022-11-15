package model

// StreamsVideoCreateRequest model
type StreamsVideoCreateRequest struct {
	// The streams input asset URL
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
}
