package model

type LiveEncoding struct {
	// Stream key of the live encoder
	StreamKey string `json:"streamKey,omitempty"`
	// IP address of the live encoder
	EncoderIp string `json:"encoderIp,omitempty"`
}

