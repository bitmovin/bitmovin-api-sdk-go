package model

// PerTitle model
type PerTitle struct {
	// Per-Title configuration for H264
	H264Configuration *H264PerTitleConfiguration `json:"h264Configuration,omitempty"`
	// Per-Title configuration for H265
	H265Configuration *H265PerTitleConfiguration `json:"h265Configuration,omitempty"`
	// Per-Title configuration for VP9
	Vp9Configuration *Vp9PerTitleConfiguration `json:"vp9Configuration,omitempty"`
	// Per-Title configuration for AV1
	Av1Configuration *Av1PerTitleConfiguration `json:"av1Configuration,omitempty"`
	// Per-Title configuration for H265 V2
	H265V2Configuration *H265V2PerTitleConfiguration `json:"h265V2Configuration,omitempty"`
}
