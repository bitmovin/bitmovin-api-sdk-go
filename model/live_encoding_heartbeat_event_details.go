package model

// LiveEncodingHeartbeatEventDetails model
type LiveEncodingHeartbeatEventDetails struct {
	EventType LiveEncodingHeartbeatEventType `json:"eventType,omitempty"`
	// Short description of the event
	Message *string `json:"message,omitempty"`
}
