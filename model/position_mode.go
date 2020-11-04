package model

// PositionMode : PositionMode model
type PositionMode string

// List of possible PositionMode values
const (
	PositionMode_KEYFRAME PositionMode = "KEYFRAME"
	PositionMode_TIME     PositionMode = "TIME"
	PositionMode_SEGMENT  PositionMode = "SEGMENT"
)
