package model

// StreamsVideoCreateRequest model
type StreamsVideoCreateRequest struct {
	// The streams input asset URL
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
	// Id of the style config to use
	StyleConfigId *string `json:"styleConfigId,omitempty"`
	// Id of the advertisement config to use
	AdConfigId *string `json:"adConfigId,omitempty"`
	// Profile to be used in encoding
	EncodingProfile StreamsEncodingProfile `json:"encodingProfile,omitempty"`
}
