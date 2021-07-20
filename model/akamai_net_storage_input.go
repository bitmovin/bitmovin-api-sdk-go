package model

import (
	"encoding/json"
)

// AkamaiNetStorageInput model
type AkamaiNetStorageInput struct {
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
	// Host to use for Akamai NetStorage transfers (required)
	Host *string `json:"host,omitempty"`
	// Your Akamai NetStorage Username (required)
	Username *string `json:"username,omitempty"`
	// Your Akamai NetStorage password (required)
	Password *string `json:"password,omitempty"`
}

func (m AkamaiNetStorageInput) InputType() InputType {
	return InputType_AKAMAI_NETSTORAGE
}
func (m AkamaiNetStorageInput) MarshalJSON() ([]byte, error) {
	type M AkamaiNetStorageInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AKAMAI_NETSTORAGE"

	return json.Marshal(x)
}
