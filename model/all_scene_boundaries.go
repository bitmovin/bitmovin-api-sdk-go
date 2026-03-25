package model

// AllSceneBoundaries model
type AllSceneBoundaries struct {
	// If true, a keyframe (IDR frame) is placed at every detected scene boundary, enabling clean segment cuts aligned with scene changes.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// If true, cue tags are inserted at every scene boundary in addition to keyframes.
	InsertCueTags *bool `json:"insertCueTags,omitempty"`
}
