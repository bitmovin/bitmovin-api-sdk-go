package model

// SimpleEncodingVodJobRequest model
type SimpleEncodingVodJobRequest struct {
	// The template that will be used for the encoding.
	EncodingTemplate EncodingTemplate                `json:"encodingTemplate,omitempty"`
	Inputs           []SimpleEncodingVodJobUrlInput  `json:"inputs,omitempty"`
	Outputs          []SimpleEncodingVodJobUrlOutput `json:"outputs,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
}
