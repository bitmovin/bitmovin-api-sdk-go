package model

import (
	"encoding/json"
)

// LocalInput model
type LocalInput struct {
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
	// Path to your local storage (required)
	Path *string `json:"path,omitempty"`
}

func (m LocalInput) InputType() InputType {
	return InputType_LOCAL
}
func (m LocalInput) MarshalJSON() ([]byte, error) {
	type M LocalInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "LOCAL"

	return json.Marshal(x)
}
