package model

// Song model
type Song struct {
	Name   *string `json:"name,omitempty"`
	Artist *string `json:"artist,omitempty"`
}
