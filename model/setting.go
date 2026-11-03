package model

// Setting model
type Setting struct {
	// Location of the scene.
	Location *Location `json:"location,omitempty"`
	// Time of day described by the analysis, in free text rather than a fixed vocabulary.
	TimeOfDay *string `json:"timeOfDay,omitempty"`
	// Description of the setting's atmosphere.
	Atmosphere *Atmosphere `json:"atmosphere,omitempty"`
}
