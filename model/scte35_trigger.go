package model

// Scte35Trigger model
type Scte35Trigger struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// The base 64 encoded data for the SCTE trigger (required)
	Base64EncodedMetadata *string `json:"base64EncodedMetadata,omitempty"`
}
