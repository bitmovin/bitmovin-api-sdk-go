package model
type AspectMode string

// List of AspectMode
const (
	AspectMode_PAD AspectMode = "PAD"
	AspectMode_CROP AspectMode = "CROP"
	AspectMode_STRETCH AspectMode = "STRETCH"
)