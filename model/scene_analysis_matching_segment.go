package model

// The scene segment that best matches a semantic-search query
type SceneAnalysisMatchingSegment struct {
	// ID of the matching scene (required)
	SceneId *string `json:"sceneId,omitempty"`
	// The detected type of the matching scene
	SceneType SceneType `json:"sceneType,omitempty"`
	// The title of the matching scene
	SceneTitle *string `json:"sceneTitle,omitempty"`
	// A description of the matching scene
	SceneDescription *string `json:"sceneDescription,omitempty"`
	// The start time of the matching segment in seconds from the beginning of the video (required)
	StartInSeconds *float64 `json:"startInSeconds,omitempty"`
	// The end time of the matching segment in seconds from the beginning of the video (required)
	EndInSeconds *float64 `json:"endInSeconds,omitempty"`
}
