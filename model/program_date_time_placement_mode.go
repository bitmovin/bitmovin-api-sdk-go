package model

// ProgramDateTimePlacementMode : Placement mode of the ProgramDateTime tag.
type ProgramDateTimePlacementMode string

// List of possible ProgramDateTimePlacementMode values
const (
	ProgramDateTimePlacementMode_ONCE_PER_PLAYLIST ProgramDateTimePlacementMode = "ONCE_PER_PLAYLIST"
	ProgramDateTimePlacementMode_SEGMENTS_INTERVAL ProgramDateTimePlacementMode = "SEGMENTS_INTERVAL"
	ProgramDateTimePlacementMode_SECONDS_INTERVAL  ProgramDateTimePlacementMode = "SECONDS_INTERVAL"
)
