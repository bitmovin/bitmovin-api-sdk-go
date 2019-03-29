package model

type EncodingStreamInput struct {
	// Input id
	InputId string `json:"inputId,omitempty"`
	// Path to media file
	InputPath string `json:"inputPath,omitempty"`
	Details *EncodingStreamInputDetails `json:"details,omitempty"`
}

