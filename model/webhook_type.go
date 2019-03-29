package model
type WebhookType string

// List of WebhookType
const (
	WebhookType_FINISHED WebhookType = "ENCODING_FINISHED"
	WebhookType_ERROR WebhookType = "ENCODING_ERROR"
)