package model

type AudioStreamDetails struct {
	Id string `json:"id,omitempty"`
	Codec string `json:"codec,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Position *int32 `json:"position,omitempty"`
	SampleRate *int32 `json:"sampleRate,omitempty"`
	Bitrate *int32 `json:"bitrate,omitempty"`
	Language string `json:"language,omitempty"`
	HearingImpaired *bool `json:"hearingImpaired,omitempty"`
}

