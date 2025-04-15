package model

// Setting model
type Setting struct {
	Location   *Location   `json:"location,omitempty"`
	TimeOfDay  *string     `json:"timeOfDay,omitempty"`
	Atmosphere *Atmosphere `json:"atmosphere,omitempty"`
}
