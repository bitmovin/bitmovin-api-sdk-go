package model

// SimpleEncodingLiveJobInput model
type SimpleEncodingLiveJobInput struct {
	// Defines which protocol is being used as input of the live stream.
	InputType SimpleEncodingLiveJobInputType `json:"inputType,omitempty"`
	// The aspect ratio of the input video stream
	InputAspectRatio SimpleEncodingLiveInputAspectRatio `json:"inputAspectRatio,omitempty"`
}
