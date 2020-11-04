package model

// DailyStatisticsPerLabel model
type DailyStatisticsPerLabel struct {
	// Date, format. yyyy-MM-dd (required)
	Date *Date `json:"date,omitempty"`
	// List of labels and their aggregated statistics (required)
	Labels []DailyStatistics `json:"labels,omitempty"`
}
