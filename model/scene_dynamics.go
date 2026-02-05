package model

// SceneDynamics model
type SceneDynamics struct {
	// The detected tension of the scene based on content analysis
	Tension SceneTension `json:"tension,omitempty"`
	// The detected pacing of the scene based on content analysis
	Pacing ScenePacing `json:"pacing,omitempty"`
}
