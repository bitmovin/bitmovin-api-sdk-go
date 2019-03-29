package model
import (
	"time"
)

type AnalyticsCountQueryRequest struct {
	// Analytics license key
	LicenseKey string `json:"licenseKey,omitempty"`
	Filters []AnalyticsFilter `json:"filters,omitempty"`
	OrderBy []AnalyticsOrderByEntry `json:"orderBy,omitempty"`
	Dimension string `json:"dimension,omitempty"`
	Interval []AnalyticsInterval `json:"interval,omitempty"`
	GroupBy []string `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 150)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
	// Start of timeframe which is queried
	Start *time.Time `json:"start,omitempty"`
	// End of timeframe which is queried
	End *time.Time `json:"end,omitempty"`
}

