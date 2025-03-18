package model

// CacheControlSettings model
type CacheControlSettings struct {
	// Cache control settings for HLS manifest.
	Hls *CacheControlSettingsHls `json:"hls,omitempty"`
	// Cache control settings for DASH manifest.
	Dash *CacheControlSettingsDash `json:"dash,omitempty"`
	// Cache control settings for segments.
	Segments *CacheControlSettingsSegments `json:"segments,omitempty"`
}
