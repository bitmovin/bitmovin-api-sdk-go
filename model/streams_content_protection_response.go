package model

// StreamsContentProtectionResponse model
type StreamsContentProtectionResponse struct {
	// The identifier of the streams content protection entity
	Id *string `json:"id,omitempty"`
	// The list of allowed domains
	AllowedDomains []string `json:"allowedDomains,omitempty"`
	// Controls if requests to content protected streams without referer header should be allowed or denied
	AllowNoReferer *bool `json:"allowNoReferer,omitempty"`
}
