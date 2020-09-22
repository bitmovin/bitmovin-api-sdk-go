package model

// CustomData model
type CustomData struct {
	// User-specific meta data. This can hold a custom JSON object.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
}
