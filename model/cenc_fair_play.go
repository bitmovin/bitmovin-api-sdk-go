package model

// CencFairPlay model
type CencFairPlay struct {
	// Initialization vector as hexadecimal string
	Iv *string `json:"iv,omitempty"`
	// URL of the licensing server. Typically starts with a skd://. Please check with your DRM provider on their required format.
	Uri *string `json:"uri,omitempty"`
}
