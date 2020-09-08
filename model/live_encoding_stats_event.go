package model
import (
	"time"
)

type LiveEncodingStatsEvent struct {
	// Timestamp of the event, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	Time *time.Time `json:"time,omitempty"`
	Details *LiveEncodingStatsEventDetails `json:"details,omitempty"`
}

