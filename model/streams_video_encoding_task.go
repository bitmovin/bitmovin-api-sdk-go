package model

// StreamsVideoEncodingTask model
type StreamsVideoEncodingTask struct {
	// Quality of the encoding
	Quality StreamsVideoQuality `json:"quality,omitempty"`
	// Current state of the encoding
	Status StreamsVideoEncodingStatus `json:"status,omitempty"`
}
