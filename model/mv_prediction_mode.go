package model

// MvPredictionMode : Motion vector prediction mode
type MvPredictionMode string

// List of possible MvPredictionMode values
const (
	MvPredictionMode_NONE     MvPredictionMode = "NONE"
	MvPredictionMode_SPATIAL  MvPredictionMode = "SPATIAL"
	MvPredictionMode_TEMPORAL MvPredictionMode = "TEMPORAL"
	MvPredictionMode_AUTO     MvPredictionMode = "AUTO"
)
