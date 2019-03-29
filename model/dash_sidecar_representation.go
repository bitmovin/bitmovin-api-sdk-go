package model

type DashSidecarRepresentation struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// UUID of an encoding
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing
	MuxingId string `json:"muxingId,omitempty"`
	// Sidecar Id
	SidecarId string `json:"sidecarId,omitempty"`
}

