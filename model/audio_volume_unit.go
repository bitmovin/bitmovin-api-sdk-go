package model

// AudioVolumeUnit : The unit in which the audio volume should be changed
type AudioVolumeUnit string

// List of possible AudioVolumeUnit values
const (
	AudioVolumeUnit_PERCENT AudioVolumeUnit = "PERCENT"
	AudioVolumeUnit_DB      AudioVolumeUnit = "DB"
)
