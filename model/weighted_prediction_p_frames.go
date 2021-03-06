package model

// WeightedPredictionPFrames : Defines the mode for weighted prediction for P-frames
type WeightedPredictionPFrames string

// List of possible WeightedPredictionPFrames values
const (
	WeightedPredictionPFrames_DISABLED WeightedPredictionPFrames = "DISABLED"
	WeightedPredictionPFrames_SIMPLE   WeightedPredictionPFrames = "SIMPLE"
	WeightedPredictionPFrames_SMART    WeightedPredictionPFrames = "SMART"
)
