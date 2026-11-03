package model

// Location model
type Location struct {
	// Short descriptive name of the place, which need not identify a real geographic location.
	Name *string `json:"name,omitempty"`
	// Free-text description of the place.
	Description *string `json:"description,omitempty"`
}
