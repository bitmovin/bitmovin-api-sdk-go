package model

// SimpleEncodingVodJobRequest model
type SimpleEncodingVodJobRequest struct {
	Inputs  []SimpleEncodingVodJobUrlInput  `json:"inputs,omitempty"`
	Outputs []SimpleEncodingVodJobUrlOutput `json:"outputs,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
}
