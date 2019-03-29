package model

type ErrorDetails struct {
	// Specific error code
	Code *int32 `json:"code,omitempty"`
	// Error group name
	Category string `json:"category,omitempty"`
	// Detailed error message
	Text string `json:"text,omitempty"`
	// Information if the encoding could potentially succeed when retrying.
	RetryHint RetryHint `json:"retryHint,omitempty"`
}

