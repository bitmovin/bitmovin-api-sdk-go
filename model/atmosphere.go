package model

// Atmosphere model
type Atmosphere struct {
	Mood     *string `json:"mood,omitempty"`
	Lighting *string `json:"lighting,omitempty"`
	Weather  *string `json:"weather,omitempty"`
}
