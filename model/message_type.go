package model


// MessageType : MessageType model
type MessageType string

// List of possible MessageType values
const (
    MessageType_ERROR MessageType = "ERROR"
    MessageType_WARNING MessageType = "WARNING"
    MessageType_INFO MessageType = "INFO"
    MessageType_DEBUG MessageType = "DEBUG"
    MessageType_TRACE MessageType = "TRACE"
)


