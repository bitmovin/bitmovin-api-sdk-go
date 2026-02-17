package model

// DnsMappingRequest model
type DnsMappingRequest struct {
	// Subdomain used for the DNS mapping. It can only contain lowercase letters, numbers and dashes (-). It can be at most 63 characters long (required)
	Subdomain *string `json:"subdomain,omitempty"`
	// Optional name for the DNS mapping
	Name *string `json:"name,omitempty"`
	// Optional description for the DNS mapping
	Description *string `json:"description,omitempty"`
}
