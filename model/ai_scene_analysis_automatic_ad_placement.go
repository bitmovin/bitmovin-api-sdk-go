package model

// AiSceneAnalysisAutomaticAdPlacement model
type AiSceneAnalysisAutomaticAdPlacement struct {
	// Ad placements schedule
	Schedule []AutomaticAdPlacementPosition `json:"schedule,omitempty"`
	// Configuration for placing keyframes and optional cue tags at every detected scene boundary.
	AllSceneBoundaries *AllSceneBoundaries `json:"allSceneBoundaries,omitempty"`
}
