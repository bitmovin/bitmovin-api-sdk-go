package model

// Person model
type Person struct {
	// Name of the credited person extracted from the credit text.
	Name *string `json:"name,omitempty"`
	// Credited role, or the portrayed character for acting credits when available.
	Role *string `json:"role,omitempty"`
	// The detected department of a person
	Department Department `json:"department,omitempty"`
}
