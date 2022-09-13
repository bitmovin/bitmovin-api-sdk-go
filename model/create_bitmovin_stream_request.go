package model

// CreateBitmovinStreamRequest model
type CreateBitmovinStreamRequest struct {
	// The Streams input asset URL
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Title of the Stream
	Title *string `json:"title,omitempty"`
	// Description of the Stream
	Description *string `json:"description,omitempty"`
}
