package model

// StreamsLiveCreateRequest model
type StreamsLiveCreateRequest struct {
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
	// Id of the domain restriction config to use
	DomainRestrictionId *string `json:"domainRestrictionId,omitempty"`
}
