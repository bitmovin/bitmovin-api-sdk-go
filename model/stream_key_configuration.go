package model

// StreamKeyConfiguration model
type StreamKeyConfiguration struct {
	Type StreamKeyConfigurationType `json:"type,omitempty"`
	// Id of the previously generated stream key.  Only needed when the type is `ASSIGN`.
	StreamKeyId *string `json:"streamKeyId,omitempty"`
}
