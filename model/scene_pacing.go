package model

// ScenePacing : ScenePacing model
type ScenePacing string

// List of possible ScenePacing values
const (
	ScenePacing_SLOW     ScenePacing = "SLOW"
	ScenePacing_MEASURED ScenePacing = "MEASURED"
	ScenePacing_BRISK    ScenePacing = "BRISK"
	ScenePacing_FRANTIC  ScenePacing = "FRANTIC"
	ScenePacing_UNKNOWN  ScenePacing = "UNKNOWN"
)
