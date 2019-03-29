package model

type VideoStreamDetails struct {
	Id string `json:"id,omitempty"`
	Codec string `json:"codec,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Fps string `json:"fps,omitempty"`
	Width *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Bitrate *int32 `json:"bitrate,omitempty"`
}

