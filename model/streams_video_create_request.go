package model

// StreamsVideoCreateRequest model
type StreamsVideoCreateRequest struct {
	// The streams input asset URL
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
	// Profile to be used in encoding
	EncodingProfile StreamsEncodingProfile `json:"encodingProfile,omitempty"`
	// If set to true the Stream is only accessible via a token
	Signed *bool `json:"signed,omitempty"`
}
