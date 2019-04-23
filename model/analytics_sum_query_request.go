package model

type AnalyticsSumQueryRequest struct {
	// Analytics license key
	LicenseKey string `json:"licenseKey,omitempty"`
	Filters []AnalyticsFilter `json:"filters,omitempty"`
	OrderBy []AnalyticsOrderByEntry `json:"orderBy,omitempty"`
	Dimension string `json:"dimension,omitempty"`
	Interval AnalyticsInterval `json:"interval,omitempty"`
	GroupBy []string `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
	// Start of timeframe which is queried
	Start string `json:"start,omitempty"`
	// End of timeframe which is queried
	End string `json:"end,omitempty"`
}

