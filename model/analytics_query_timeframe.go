package model

// AnalyticsQueryTimeframe model
type AnalyticsQueryTimeframe struct {
	// Start of timeframe which is queried in UTC format.
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *DateTime `json:"end,omitempty"`
}
