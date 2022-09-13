package model

// UpdateBitmovinStreamRequest model
type UpdateBitmovinStreamRequest struct {
	// The status of the Stream
	Status BitmovinStreamStatus `json:"status,omitempty"`
	// Title of the Stream
	Title *string `json:"title,omitempty"`
	// Description of the Stream
	Description *string `json:"description,omitempty"`
}
