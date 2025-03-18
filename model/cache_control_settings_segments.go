package model

// CacheControlSettingsSegments model
type CacheControlSettingsSegments struct {
	// Cache control settings for init segment.
	InitSegment *CacheControl `json:"initSegment,omitempty"`
	// Cache control settings for media segment.
	MediaSegment *CacheControl `json:"mediaSegment,omitempty"`
}
