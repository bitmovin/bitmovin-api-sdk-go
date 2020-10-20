package model

// EnhancedDeinterlaceMode : Specifies the method how fields are converted to frames
type EnhancedDeinterlaceMode string

// List of possible EnhancedDeinterlaceMode values
const (
	EnhancedDeinterlaceMode_FRAME EnhancedDeinterlaceMode = "FRAME"
	EnhancedDeinterlaceMode_FIELD EnhancedDeinterlaceMode = "FIELD"
)
