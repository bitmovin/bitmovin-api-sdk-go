package model

// SimpleEncodingVodJobErrors model
type SimpleEncodingVodJobErrors struct {
	// Stable code that identifies the error type.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Human readable description of the error.
	Message *string `json:"message,omitempty"`
	// Additional details of the error if available.
	Details *string `json:"details,omitempty"`
}
