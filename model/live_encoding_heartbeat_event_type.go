package model

// LiveEncodingHeartbeatEventType : LiveEncodingHeartbeatEventType model
type LiveEncodingHeartbeatEventType string

// List of possible LiveEncodingHeartbeatEventType values
const (
	LiveEncodingHeartbeatEventType_FIRST_CONNECT LiveEncodingHeartbeatEventType = "FIRST_CONNECT"
	LiveEncodingHeartbeatEventType_DISCONNECT    LiveEncodingHeartbeatEventType = "DISCONNECT"
	LiveEncodingHeartbeatEventType_RECONNECT     LiveEncodingHeartbeatEventType = "RECONNECT"
	LiveEncodingHeartbeatEventType_WARNING       LiveEncodingHeartbeatEventType = "WARNING"
	LiveEncodingHeartbeatEventType_ERROR         LiveEncodingHeartbeatEventType = "ERROR"
)
