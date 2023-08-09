package model

// StreamsSearchResponse model
type StreamsSearchResponse struct {
	// The identifier of the stream
	Id *string `json:"id,omitempty"`
	// The title of the stream
	Title *string `json:"title,omitempty"`
	// The description of the stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// The type the stream
	Type StreamsType `json:"type,omitempty"`
}
