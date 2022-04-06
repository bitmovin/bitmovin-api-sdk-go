package model

import (
	"encoding/json"
)

// CdnOutput model
type CdnOutput struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource
	Name *string `json:"name,omitempty"`
	// Description for the resource
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	Acl        []AclEntry              `json:"acl,omitempty"`
	// Domain name for public access to CDN content
	DomainName *string `json:"domainName,omitempty"`
}

func (m CdnOutput) OutputType() OutputType {
	return OutputType_CDN
}
func (m CdnOutput) MarshalJSON() ([]byte, error) {
	type M CdnOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CDN"

	return json.Marshal(x)
}
