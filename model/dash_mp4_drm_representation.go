package model

type DashMp4DrmRepresentation struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// UUID of an encoding
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing
	MuxingId string `json:"muxingId,omitempty"`
	// Path to the MP4 file
	FilePath string `json:"filePath,omitempty"`
	// DRM Id
	DrmId string `json:"drmId,omitempty"`
}

