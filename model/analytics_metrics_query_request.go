package model

type AnalyticsMetricsQueryRequest struct {
	// Start of timeframe which is queried
	Start string `json:"start,omitempty"`
	// End of timeframe which is queried
	End string `json:"end,omitempty"`
	// Analytics license key (required)
	LicenseKey string `json:"licenseKey,omitempty"`
	Filters []AnalyticsFilter `json:"filters,omitempty"`
	OrderBy []AnalyticsOrderByEntry `json:"orderBy,omitempty"`
	Interval AnalyticsInterval `json:"interval,omitempty"`
	GroupBy []AnalyticsAttribute `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data used for pagination
	Offset *int64 `json:"offset,omitempty"`
}

