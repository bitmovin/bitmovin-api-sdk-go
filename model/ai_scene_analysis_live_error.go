package model

// AiSceneAnalysisLiveError model
type AiSceneAnalysisLiveError struct {
	// Stable machine-readable failure code (required)
	Code *string `json:"code,omitempty"`
	// Credential-free failure description safe to expose to the customer (required)
	Message *string `json:"message,omitempty"`
	// Time at which the failure was recorded (required)
	Timestamp *DateTime `json:"timestamp,omitempty"`
}
