package model

// StreamsLiveResponse model
type StreamsLiveResponse struct {
	// The identifier of the stream
	Id *string `json:"id,omitempty"`
	// The streamKey of the stream
	StreamKey *string `json:"streamKey,omitempty"`
	// The title of the stream
	Title *string `json:"title,omitempty"`
	// The description of the stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// The life cycle of the stream
	LifeCycle StreamsLiveLifeCycle   `json:"lifeCycle,omitempty"`
	Config    *StreamsConfigResponse `json:"config,omitempty"`
	// Poster URL
	PosterUrl *string `json:"posterUrl,omitempty"`
}
