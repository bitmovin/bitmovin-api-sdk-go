package model

// AnalyticsOrderByEntry model
type AnalyticsOrderByEntry struct {
	Name  AnalyticsAttribute `json:"name,omitempty"`
	Order AnalyticsOrder     `json:"order,omitempty"`
}
