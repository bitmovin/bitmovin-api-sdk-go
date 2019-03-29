package model
import (
	"time"
)

type DailyStatisticsPerLabel struct {
	// Date, format. yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`
	// List of labels and their aggregated statistics
	Labels []DailyStatistics `json:"labels,omitempty"`
}

