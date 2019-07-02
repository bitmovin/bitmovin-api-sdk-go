package model

type DashRepresentation struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId string `json:"muxingId,omitempty"`
}

