package model

type StreamPerTitleFixedResolutionAndBitrateSettings struct {
	// The minimum bitrate that will be used for that template.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// The maximum bitrate that will be used for that template.
	MaxBitrate *int32 `json:"maxBitrate,omitempty"`
	// Bitrate selection mode
	BitrateSelectionMode BitrateSelectionMode `json:"bitrateSelectionMode,omitempty"`
	// Low complexity boundary for max bitrate
	LowComplexityBoundaryForMaxBitrate *float64 `json:"lowComplexityBoundaryForMaxBitrate,omitempty"`
	// High complexity boundary for max bitrate
	HighComplexityBoundaryForMaxBitrate *float64 `json:"highComplexityBoundaryForMaxBitrate,omitempty"`
}

