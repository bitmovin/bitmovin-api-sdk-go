package model

// DnsMappingResponse model
type DnsMappingResponse struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Subdomain used for the DNS mapping
	Subdomain *string `json:"subdomain,omitempty"`
	// Resolved hostname for the DNS mapping
	Hostname *string `json:"hostname,omitempty"`
	// ID of the encoding this DNS mapping belongs to
	EncodingId *string `json:"encodingId,omitempty"`
	// IP address that the hostname resolves to
	Ip *string `json:"ip,omitempty"`
}
