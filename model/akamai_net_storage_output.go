package model

import (
	"encoding/json"
)

// AkamaiNetStorageOutput model
type AkamaiNetStorageOutput struct {
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
	// Deprecation notice: This property does not have any effect and will not be returned by GET endpoints
	Acl []AclEntry `json:"acl,omitempty"`
	// Host to use for Akamai NetStorage transfers (required)
	Host *string `json:"host,omitempty"`
	// Your Akamai NetStorage Username (required)
	Username *string `json:"username,omitempty"`
	// Your Akamai NetStorage password (required)
	Password *string `json:"password,omitempty"`
}

func (m AkamaiNetStorageOutput) OutputType() OutputType {
	return OutputType_AKAMAI_NETSTORAGE
}
func (m AkamaiNetStorageOutput) MarshalJSON() ([]byte, error) {
	type M AkamaiNetStorageOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AKAMAI_NETSTORAGE"

	return json.Marshal(x)
}
