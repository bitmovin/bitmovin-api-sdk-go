package model

// BitmovinStreamEncodingTask model
type BitmovinStreamEncodingTask struct {
	// Quality of the encoding
	Quality BitmovinStreamQuality `json:"quality,omitempty"`
	// Current state of the encoding
	Status BitmovinStreamEncodingStatus `json:"status,omitempty"`
}
