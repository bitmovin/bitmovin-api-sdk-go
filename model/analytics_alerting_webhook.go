package model

// AnalyticsAlertingWebhook model
type AnalyticsAlertingWebhook struct {
	// Webhook url (required)
	Url *string `json:"url,omitempty"`
	// HTTP method used for the webhook (required)
	Method WebhookHttpMethod `json:"method,omitempty"`
	// Whether to skip SSL certification verification or not (required)
	InsecureSsl *bool `json:"insecureSsl,omitempty"`
}
