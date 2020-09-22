package model

// StreamInfos model
type StreamInfos struct {
	// Timestamp of the event, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	Time *DateTime `json:"time,omitempty"`
	// Details about billable minutes for each resolution category
	StreamInfos []StreamInfosDetails `json:"streamInfos,omitempty"`
}
