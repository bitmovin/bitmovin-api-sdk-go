package model

// Vp9Quality : Determines quality for the price of speed.
type Vp9Quality string

// List of possible Vp9Quality values
const (
	Vp9Quality_REALTIME Vp9Quality = "REALTIME"
	Vp9Quality_GOOD     Vp9Quality = "GOOD"
	Vp9Quality_BEST     Vp9Quality = "BEST"
)
