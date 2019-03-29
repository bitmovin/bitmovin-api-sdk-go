package model

type SubtitleStreamDetails struct {
	Id string `json:"id,omitempty"`
	Codec string `json:"codec,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Language string `json:"language,omitempty"`
	HearingImpaired *bool `json:"hearingImpaired,omitempty"`
}

