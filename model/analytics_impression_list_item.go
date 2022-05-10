package model

// AnalyticsImpressionListItem model
type AnalyticsImpressionListItem struct {
	// Random UUID that is used to identify a session (required)
	ImpressionId *string `json:"impressionId,omitempty"`
}
