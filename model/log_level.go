package model


// LogLevel : LogLevel model
type LogLevel string

// List of possible LogLevel values
const (
    LogLevel_TRACE LogLevel = "TRACE"
    LogLevel_DEBUG LogLevel = "DEBUG"
    LogLevel_INFO LogLevel = "INFO"
    LogLevel_WARN LogLevel = "WARN"
    LogLevel_ERROR LogLevel = "ERROR"
    LogLevel_FATAL LogLevel = "FATAL"
    LogLevel_OFF LogLevel = "OFF"
)


