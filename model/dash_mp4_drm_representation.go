package model

type DashMp4DrmRepresentation struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId string `json:"muxingId,omitempty"`
	// Path to the MP4 file (required)
	FilePath string `json:"filePath,omitempty"`
	// DRM Id (required)
	DrmId string `json:"drmId,omitempty"`
}

