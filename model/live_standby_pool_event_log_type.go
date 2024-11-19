package model

// LiveStandbyPoolEventLogType : Event log type of the standby pool
type LiveStandbyPoolEventLogType string

// List of possible LiveStandbyPoolEventLogType values
const (
	LiveStandbyPoolEventLogType_INFO  LiveStandbyPoolEventLogType = "INFO"
	LiveStandbyPoolEventLogType_WARN  LiveStandbyPoolEventLogType = "WARN"
	LiveStandbyPoolEventLogType_ERROR LiveStandbyPoolEventLogType = "ERROR"
)
