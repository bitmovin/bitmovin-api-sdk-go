package model

// Content model
type Content struct {
	Characters []Character   `json:"characters,omitempty"`
	Objects    []SceneObject `json:"objects,omitempty"`
	Settings   []Setting     `json:"settings,omitempty"`
}
