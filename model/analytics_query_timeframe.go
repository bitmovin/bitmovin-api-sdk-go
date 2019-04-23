package model

type AnalyticsQueryTimeframe struct {
	// Start of timeframe which is queried
	Start string `json:"start,omitempty"`
	// End of timeframe which is queried
	End string `json:"end,omitempty"`
}

