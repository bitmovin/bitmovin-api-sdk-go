package model

// EnhancedDeinterlaceAutoEnable : Specifies if the Enhanced Deinterlace filter should be applied unconditionally or only on demand.
type EnhancedDeinterlaceAutoEnable string

// List of possible EnhancedDeinterlaceAutoEnable values
const (
	EnhancedDeinterlaceAutoEnable_ALWAYS_ON       EnhancedDeinterlaceAutoEnable = "ALWAYS_ON"
	EnhancedDeinterlaceAutoEnable_META_DATA_BASED EnhancedDeinterlaceAutoEnable = "META_DATA_BASED"
)
