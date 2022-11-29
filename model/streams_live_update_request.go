package model

// StreamsLiveUpdateRequest model
type StreamsLiveUpdateRequest struct {
	// The new title of the stream
	Title *string `json:"title,omitempty"`
	// The new description of the stream
	Description *string `json:"description,omitempty"`
}
