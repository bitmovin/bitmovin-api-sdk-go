package model

// StreamKeyStatus : Status of the stream key
type StreamKeyStatus string

// List of possible StreamKeyStatus values
const (
	StreamKeyStatus_ASSIGNED   StreamKeyStatus = "ASSIGNED"
	StreamKeyStatus_UNASSIGNED StreamKeyStatus = "UNASSIGNED"
)
