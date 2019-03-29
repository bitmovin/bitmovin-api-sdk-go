package model

type AudioGroup struct {
	// Name of the audio group
	Name string `json:"name,omitempty"`
	// Priority of the audio group
	Priority *int32 `json:"priority,omitempty"`
}

