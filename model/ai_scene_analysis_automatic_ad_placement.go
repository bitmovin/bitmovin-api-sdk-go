package model

// AiSceneAnalysisAutomaticAdPlacement model
type AiSceneAnalysisAutomaticAdPlacement struct {
	// Ad placements schedule
	Schedule []AutomaticAdPlacementPosition `json:"schedule,omitempty"`
}
