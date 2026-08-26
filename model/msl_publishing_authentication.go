package model

// MSL5 publishing-side authentication. When enabled, the encoder sends HTTP Digest Authentication headers with every segment upload.  When `enabled` is `true`, `username` and `password` are required; the API rejects the request otherwise. When `enabled` is `false` (or this object is omitted), credentials are ignored.
type MslPublishingAuthentication struct {
	// Whether HTTP Digest publishing authentication is enabled. (required)
	Enabled *bool `json:"enabled,omitempty"`
	// HTTP Digest username for publishing MSL5 segments. Required when `enabled` is `true`.
	Username *string `json:"username,omitempty"`
	// HTTP Digest password for publishing MSL5 segments. Required when `enabled` is `true`.
	Password *string `json:"password,omitempty"`
}
