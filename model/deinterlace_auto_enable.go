package model

// DeinterlaceAutoEnable : Specifies if the Deinterlace Filter should be applied unconditionally or only on demand.
type DeinterlaceAutoEnable string

// List of possible DeinterlaceAutoEnable values
const (
	DeinterlaceAutoEnable_ALWAYS_ON                   DeinterlaceAutoEnable = "ALWAYS_ON"
	DeinterlaceAutoEnable_META_DATA_BASED             DeinterlaceAutoEnable = "META_DATA_BASED"
	DeinterlaceAutoEnable_META_DATA_AND_CONTENT_BASED DeinterlaceAutoEnable = "META_DATA_AND_CONTENT_BASED"
)
