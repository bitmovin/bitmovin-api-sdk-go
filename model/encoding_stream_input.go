package model

type EncodingStreamInput struct {
	// Input id (required)
	InputId string `json:"inputId,omitempty"`
	// Path to media file (required)
	InputPath string `json:"inputPath,omitempty"`
	Details *EncodingStreamInputDetails `json:"details,omitempty"`
}

