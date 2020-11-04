package model

// MaxCtuSize : Set the maximal CTU (Coding Tree Unit) size
type MaxCtuSize string

// List of possible MaxCtuSize values
const (
	MaxCtuSize_S16 MaxCtuSize = "16"
	MaxCtuSize_S32 MaxCtuSize = "32"
	MaxCtuSize_S64 MaxCtuSize = "64"
)
