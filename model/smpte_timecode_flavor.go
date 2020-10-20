package model

// SmpteTimecodeFlavor : SmpteTimecodeFlavor model
type SmpteTimecodeFlavor string

// List of possible SmpteTimecodeFlavor values
const (
	SmpteTimecodeFlavor_AUTO           SmpteTimecodeFlavor = "AUTO"
	SmpteTimecodeFlavor_NON_DROP_FRAME SmpteTimecodeFlavor = "NON_DROP_FRAME"
	SmpteTimecodeFlavor_DROP_FRAME     SmpteTimecodeFlavor = "DROP_FRAME"
)
