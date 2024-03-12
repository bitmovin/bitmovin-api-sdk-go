package model

// StreamsVideoUpdateRequest model
type StreamsVideoUpdateRequest struct {
	// The new status of the stream
	Status StreamsVideoStatus `json:"status,omitempty"`
	// The new title of the stream
	Title *string `json:"title,omitempty"`
	// The new description of the stream
	Description *string `json:"description,omitempty"`
	// URL to hosted poster image
	PosterUrl *string `json:"posterUrl,omitempty"`
	// Id of the domain restriction config to use
	DomainRestrictionId *string `json:"domainRestrictionId,omitempty"`
}
