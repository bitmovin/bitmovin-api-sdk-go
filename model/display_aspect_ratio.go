package model

type DisplayAspectRatio struct {
	// The numerator of the display aspect ratio (DAR). For example for a DAR of 16:9, the value 16 must be used. (required)
	Numerator *int32 `json:"numerator,omitempty"`
	// The denominator of the display aspect ratio (DAR). For example for a DAR of 16:9, the value 9 must be used. (required)
	Denominator *int32 `json:"denominator,omitempty"`
}

