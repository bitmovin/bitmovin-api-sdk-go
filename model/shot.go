package model

// A detected shot associated with one scene. A shot can cross a scene boundary or lie entirely outside its associated scene. Use the shot's own timestamps to locate its content.
type Shot struct {
	// The start time of the shot in seconds from the beginning of the video (required)
	StartInSeconds *float64 `json:"startInSeconds,omitempty"`
	// The end time of the shot in seconds from the beginning of the video (required)
	EndInSeconds *float64 `json:"endInSeconds,omitempty"`
	// Description of the shot's visual content, action and context.
	DetailedDescription *string `json:"detailedDescription,omitempty"`
	// Keywords describing the shot's content, themes or notable elements.
	Keywords []string `json:"keywords,omitempty"`
	// Selected subjects or objects identified in this shot.
	MainSubjects []MainSubject `json:"mainSubjects,omitempty"`
	// Content advisory results for this shot, or null when no assessment is available. Availability depends on whether advisory analysis was enabled and supported by the analysis version.
	ShotAdvisories *ShotAdvisories `json:"shotAdvisories,omitempty"`
}
