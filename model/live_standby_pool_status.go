package model

// LiveStandbyPoolStatus : Status of the live standby pool
type LiveStandbyPoolStatus string

// List of possible LiveStandbyPoolStatus values
const (
	LiveStandbyPoolStatus_HEALTHY LiveStandbyPoolStatus = "HEALTHY"
	LiveStandbyPoolStatus_ERROR   LiveStandbyPoolStatus = "ERROR"
)
