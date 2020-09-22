package model

// AdaptiveQuantMode : AdaptiveQuantMode model
type AdaptiveQuantMode string

// List of possible AdaptiveQuantMode values
const (
	AdaptiveQuantMode_DISABLED                  AdaptiveQuantMode = "DISABLED"
	AdaptiveQuantMode_VARIANCE                  AdaptiveQuantMode = "VARIANCE"
	AdaptiveQuantMode_AUTO_VARIANCE             AdaptiveQuantMode = "AUTO_VARIANCE"
	AdaptiveQuantMode_AUTO_VARIANCE_DARK_SCENES AdaptiveQuantMode = "AUTO_VARIANCE_DARK_SCENES"
)
