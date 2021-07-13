package model

// ImscStylingMode : Determines how IMSC styling is handled.
type ImscStylingMode string

// List of possible ImscStylingMode values
const (
	ImscStylingMode_PASSTHROUGH  ImscStylingMode = "PASSTHROUGH"
	ImscStylingMode_DROP_STYLING ImscStylingMode = "DROP_STYLING"
)
