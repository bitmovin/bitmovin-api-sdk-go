package model

// LiveEncodingStatsEventDetails model
type LiveEncodingStatsEventDetails struct {
	EventType LiveEncodingEventName `json:"eventType,omitempty"`
	// Short description of the event
	Message *string `json:"message,omitempty"`
	// Additional event details as key-value pairs
	AdditionalProperties *string `json:"additionalProperties,omitempty"`
}
