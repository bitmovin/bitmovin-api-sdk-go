package model

// SimpleEncodingLiveJobInput model
type SimpleEncodingLiveJobInput struct {
	// Defines which protocol is being used as input of the live stream.
	InputType SimpleEncodingLiveJobInputType `json:"inputType,omitempty"`
}
