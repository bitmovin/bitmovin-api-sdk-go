package model

// CacheControlSettingsDash model
type CacheControlSettingsDash struct {
	// Cache control settings for DASH Timeline manifest.
	TimelineManifest *CacheControl `json:"timelineManifest,omitempty"`
	// Cache control settings for DASH Template manifest.
	TemplateManifest *CacheControl `json:"templateManifest,omitempty"`
}
