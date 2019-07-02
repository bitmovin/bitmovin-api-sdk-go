package model

type LiveEncoding struct {
	// Stream key of the live encoder (required)
	StreamKey string `json:"streamKey,omitempty"`
	// IP address of the live encoder (required)
	EncoderIp string `json:"encoderIp,omitempty"`
}

