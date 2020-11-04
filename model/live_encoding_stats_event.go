package model

// LiveEncodingStatsEvent model
type LiveEncodingStatsEvent struct {
	// Timestamp of the event, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	Time    *DateTime                      `json:"time,omitempty"`
	Details *LiveEncodingStatsEventDetails `json:"details,omitempty"`
}
