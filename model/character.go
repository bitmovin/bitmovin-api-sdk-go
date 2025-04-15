package model

// Character model
type Character struct {
	Appearance  *string `json:"appearance,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
