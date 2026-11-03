package model

// SceneObject model
type SceneObject struct {
	// Free-text description of the object in this scene.
	Description *string `json:"description,omitempty"`
	// Kind of object in free text rather than a fixed category list.
	Category *string `json:"category,omitempty"`
}
