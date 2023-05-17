package model

// StreamsLiveCreateRequest model
type StreamsLiveCreateRequest struct {
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
	// Id of the stream config to use
	ConfigId *string `json:"configId,omitempty"`
	// Id of the advertisement config to use
	AdConfigId *string `json:"adConfigId,omitempty"`
}
