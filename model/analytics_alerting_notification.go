package model

// AnalyticsAlertingNotification model
type AnalyticsAlertingNotification struct {
	// List of email addresses
	Emails   []string                   `json:"emails,omitempty"`
	Webhooks []AnalyticsAlertingWebhook `json:"webhooks,omitempty"`
}
