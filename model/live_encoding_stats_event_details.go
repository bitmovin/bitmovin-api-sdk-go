package model

// LiveEncodingStatsEventDetails model
type LiveEncodingStatsEventDetails struct {
	EventType LiveEncodingEventName `json:"eventType,omitempty"`
	// Short description of the event
	Message *string `json:"message,omitempty"`
	// Source used for clock-synchronization
	Source ClockSynchronizationMode `json:"source,omitempty"`
	// Year specified in picture timing
	Year *int64 `json:"year,omitempty"`
	// Month specified in picture timing
	Month *int64 `json:"month,omitempty"`
	// Day specified in picture timing
	Day *int64 `json:"day,omitempty"`
	// Hours specified in picture timing
	Hours *int64 `json:"hours,omitempty"`
	// Minutes specified in picture timing
	Minutes *int64 `json:"minutes,omitempty"`
	// Seconds specified in picture timing
	Seconds *int64 `json:"seconds,omitempty"`
	// Microseconds specified in picture timing
	MicroSeconds *int64 `json:"microSeconds,omitempty"`
}
