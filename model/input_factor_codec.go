package model

type InputFactorCodec struct {
	Value string `json:"value,omitempty"`
	Factor *float32 `json:"factor,omitempty"`
}

