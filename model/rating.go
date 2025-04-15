package model

// Rating model
type Rating struct {
	Region *string `json:"region,omitempty"`
	System *string `json:"system,omitempty"`
	Rating *string `json:"rating,omitempty"`
}
