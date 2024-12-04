package model

// LiveEncodingStatsEventDetails model
type LiveEncodingStatsEventDetails struct {
	EventType LiveEncodingEventName `json:"eventType,omitempty"`
	// Short description of the event
	Message *string `json:"message,omitempty"`
	// Name of the mid roll asset name
	MidRollAssetNames []string `json:"midRollAssetNames,omitempty"`
	// Duration in seconds
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty"`
}
