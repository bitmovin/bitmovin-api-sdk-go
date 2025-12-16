package model

// ClockSynchronizationMode : Mode of clock synchronization for ad insertion
type ClockSynchronizationMode string

// List of possible ClockSynchronizationMode values
const (
	ClockSynchronizationMode_SYSTEM_CLOCK ClockSynchronizationMode = "SYSTEM_CLOCK"
	ClockSynchronizationMode_EMBEDDED     ClockSynchronizationMode = "EMBEDDED"
)
