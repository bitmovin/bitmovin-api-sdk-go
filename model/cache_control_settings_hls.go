package model

// CacheControlSettingsHls model
type CacheControlSettingsHls struct {
	// Cache control settings for HLS Multivariant playlist.
	MultiVariantPlaylist *CacheControl `json:"multiVariantPlaylist,omitempty"`
	// Cache control settings for HLS Media playlist.
	VariantPlaylist *CacheControl `json:"variantPlaylist,omitempty"`
}
