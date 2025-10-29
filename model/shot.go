package model

// Represents a continuous camera shot within a scene, containing detailed visual analysis including subjects, timing, and descriptive metadata
type Shot struct {
	// The start time of the shot in seconds from the beginning of the video (required)
	StartInSeconds *float32 `json:"startInSeconds,omitempty"`
	// The end time of the shot in seconds from the beginning of the video (required)
	EndInSeconds *float32 `json:"endInSeconds,omitempty"`
	// A comprehensive textual description of the visual content, action, and context within this shot
	DetailedDescription *string `json:"detailedDescription,omitempty"`
	// A list of relevant keywords and tags that describe the content, themes, or notable elements in this shot
	Keywords []string `json:"keywords,omitempty"`
	// A collection of the primary subjects or objects detected and tracked within this shot, including their positions and characteristics
	MainSubjects []MainSubject `json:"mainSubjects,omitempty"`
}
