package model

// CacheControl model
type CacheControl struct {
	// Cache control for storing data on CDN. Example \"public, max-age=0, no-cache\".
	CacheControl *string `json:"cacheControl,omitempty"`
}
