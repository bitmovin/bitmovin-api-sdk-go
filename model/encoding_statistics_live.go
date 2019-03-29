package model
import (
	"time"
)

type EncodingStatisticsLive struct {
	// Date, format. yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`
	// Bytes encoded for this encoding.
	BytesEncoded *int64 `json:"bytesEncoded,omitempty"`
	// Time in seconds encoded for this encoding.
	TimeEncoded *int64 `json:"timeEncoded,omitempty"`
}

