package model

// EsamDirection : ESAM direction indicator following the SCTE-250 standard
type EsamDirection string

// List of possible EsamDirection values
const (
	EsamDirection_OUT  EsamDirection = "OUT"
	EsamDirection_IN   EsamDirection = "IN"
	EsamDirection_BOTH EsamDirection = "BOTH"
)
