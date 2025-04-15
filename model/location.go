package model

// Location model
type Location struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
