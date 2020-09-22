package model

// DeinterlaceFrameSelectionMode : Specifies which frames to deinterlace
type DeinterlaceFrameSelectionMode string

// List of possible DeinterlaceFrameSelectionMode values
const (
	DeinterlaceFrameSelectionMode_ALL        DeinterlaceFrameSelectionMode = "ALL"
	DeinterlaceFrameSelectionMode_INTERLACED DeinterlaceFrameSelectionMode = "INTERLACED"
)
