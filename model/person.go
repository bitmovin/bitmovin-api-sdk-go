package model

// Person model
type Person struct {
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
	// The detected department of a person
	Department Department `json:"department,omitempty"`
}
