package model

// LiveStandbyPoolEncodingIngestPoint model
type LiveStandbyPoolEncodingIngestPoint struct {
	// URL to the RTMP/RTMPS endpoint for this live encoding
	StreamBaseUrl *string `json:"streamBaseUrl,omitempty"`
	// Stream key value of this live encoding
	StreamKey *string `json:"streamKey,omitempty"`
}
