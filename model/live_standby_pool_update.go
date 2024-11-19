package model

// LiveStandbyPoolUpdate model
type LiveStandbyPoolUpdate struct {
	// Number of instances to keep ready for streaming while the pool is running
	TargetPoolSize *int32 `json:"targetPoolSize,omitempty"`
	// Base64 encoded template used to start the encodings in the pool
	EncodingTemplate *string `json:"encodingTemplate,omitempty"`
}
