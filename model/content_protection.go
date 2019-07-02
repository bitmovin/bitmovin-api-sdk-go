package model

type ContentProtection struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId string `json:"muxingId,omitempty"`
	// DRM Id (required)
	DrmId string `json:"drmId,omitempty"`
}

