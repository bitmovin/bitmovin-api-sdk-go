package model

type AnalyticsOrderByEntry struct {
	Name string `json:"name,omitempty"`
	Order AnalyticsOrder `json:"order,omitempty"`
}

