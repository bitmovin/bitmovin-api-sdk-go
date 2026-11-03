package model

// Atmosphere model
type Atmosphere struct {
	// Mood or tone of the setting, in free text.
	Mood *string `json:"mood,omitempty"`
	// Description of the lighting in free text.
	Lighting *string `json:"lighting,omitempty"`
	// Weather or indoor conditions described by the analysis, in free text.
	Weather *string `json:"weather,omitempty"`
}
