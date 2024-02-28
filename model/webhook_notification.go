package model

// WebhookNotification model
type WebhookNotification struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Notify when condition resolves after it was met
	Resolve *bool `json:"resolve,omitempty"`
	// Specific resource, e.g. encoding id
	ResourceId *string `json:"resourceId,omitempty"`
	// Last time the notification was triggered
	TriggeredAt  *DateTime `json:"triggeredAt,omitempty"`
	Type         *string   `json:"type,omitempty"`
	EventType    *string   `json:"eventType,omitempty"`
	Category     *string   `json:"category,omitempty"`
	ResourceType *string   `json:"resourceType,omitempty"`
	Muted        *bool     `json:"muted,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// The destination URL where the webhook data is send to (required)
	Url *string `json:"url,omitempty"`
	// HTTP method used for the webhook
	Method WebhookHttpMethod `json:"method,omitempty"`
	// Skip verification of the SSL certificate
	InsecureSsl *bool `json:"insecureSsl,omitempty"`
	// Signature used for the webhook
	Signature *WebhookSignature `json:"signature,omitempty"`
}
