package model

// LiveEncodingHeartbeatEvent model
type LiveEncodingHeartbeatEvent struct {
	// Timestamp of the event, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	Time    *DateTime                          `json:"time,omitempty"`
	Details *LiveEncodingHeartbeatEventDetails `json:"details,omitempty"`
}
