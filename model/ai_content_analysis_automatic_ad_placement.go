package model

// AiContentAnalysisAutomaticAdPlacement model
type AiContentAnalysisAutomaticAdPlacement struct {
	// Ad placements schedule
	Schedule []AutomaticAdPlacementPosition `json:"schedule,omitempty"`
}
