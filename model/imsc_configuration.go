package model

import (
	"encoding/json"
)

// ImscConfiguration model
type ImscConfiguration struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user. (required)
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	Styling    *ImscStyling            `json:"styling,omitempty"`
}

func (m ImscConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_IMSC
}
func (m ImscConfiguration) MarshalJSON() ([]byte, error) {
	type M ImscConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "IMSC"

	return json.Marshal(x)
}
