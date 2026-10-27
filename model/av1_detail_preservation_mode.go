package model

// Av1DetailPreservationMode : Controls encoder tuning that preserves film grain and fine detail at a negligible bitrate increase
type Av1DetailPreservationMode string

// List of possible Av1DetailPreservationMode values
const (
	Av1DetailPreservationMode_AUTO Av1DetailPreservationMode = "AUTO"
	Av1DetailPreservationMode_ON   Av1DetailPreservationMode = "ON"
	Av1DetailPreservationMode_OFF  Av1DetailPreservationMode = "OFF"
)
