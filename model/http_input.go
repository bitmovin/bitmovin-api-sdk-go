package model

import (
	"encoding/json"
)

// HttpInput model
type HttpInput struct {
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
	// Host Url or IP of the HTTP server (required)
	Host *string `json:"host,omitempty"`
	// Basic Auth Username, if required
	Username *string `json:"username,omitempty"`
	// Basic Auth Password, if required
	Password *string `json:"password,omitempty"`
	// Custom Port
	Port *int32 `json:"port,omitempty"`
}

func (m HttpInput) InputType() InputType {
	return InputType_HTTP
}
func (m HttpInput) MarshalJSON() ([]byte, error) {
	type M HttpInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "HTTP"

	return json.Marshal(x)
}
