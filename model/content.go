package model

// Content model
type Content struct {
	// Selected people or characters described in this scene. An empty or omitted list can mean that no characters were identified or that character descriptions were unavailable. (required)
	Characters []Character `json:"characters,omitempty"`
	// Selected notable objects described in this scene; not an exhaustive object inventory. (required)
	Objects []SceneObject `json:"objects,omitempty"`
	// Places and environmental conditions described in this scene. More than one setting can be reported. (required)
	Settings []Setting `json:"settings,omitempty"`
	// Narrative dynamics of this scene.
	Dynamics *SceneDynamics `json:"dynamics,omitempty"`
}
