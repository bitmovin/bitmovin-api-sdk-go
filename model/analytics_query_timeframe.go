package model
import (
	"time"
)

type AnalyticsQueryTimeframe struct {
	// Start of timeframe which is queried
	Start *time.Time `json:"start,omitempty"`
	// End of timeframe which is queried
	End *time.Time `json:"end,omitempty"`
}

