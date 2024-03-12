package model

// StreamsDomainRestrictionResponse model
type StreamsDomainRestrictionResponse struct {
	// The identifier of the streams domain restriction entity
	Id *string `json:"id,omitempty"`
	// The list of allowed domains
	AllowedDomains []string `json:"allowedDomains,omitempty"`
	// Controls if requests to domain restricted streams without referer header should be allowed or denied
	AllowNoReferer *bool `json:"allowNoReferer,omitempty"`
	// Controls if Stream is accessible via sharing URL or not
	AllowShare *bool `json:"allowShare,omitempty"`
}
