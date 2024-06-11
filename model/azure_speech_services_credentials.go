package model

// AzureSpeechServicesCredentials model
type AzureSpeechServicesCredentials struct {
	// Azure Speech Services resource key (required)
	SubscriptionKey *string `json:"subscriptionKey,omitempty"`
}
