package model

// StaticIpStatus : StaticIpStatus model
type StaticIpStatus string

// List of possible StaticIpStatus values
const (
	StaticIpStatus_CREATING StaticIpStatus = "CREATING"
	StaticIpStatus_UNUSED   StaticIpStatus = "UNUSED"
	StaticIpStatus_ERROR    StaticIpStatus = "ERROR"
	StaticIpStatus_USED     StaticIpStatus = "USED"
)
