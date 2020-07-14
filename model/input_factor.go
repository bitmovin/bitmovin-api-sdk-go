package model

type InputFactor struct {
	Codec *InputFactorCodec `json:"codec,omitempty"`
	Bitrate *InputFactorBitrate `json:"bitrate,omitempty"`
}

