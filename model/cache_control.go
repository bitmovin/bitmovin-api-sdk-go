package model

// CacheControl model
type CacheControl struct {
	// Cache control for storing data on CDN. Example \"public, max-age=0, no-cache\". Cache control is supported on S3, GCS and Azure output storage providers.
	CacheControl *string `json:"cacheControl,omitempty"`
}
