package model

import (
	"encoding/json"
)

// MarlinDrm model
type MarlinDrm struct {
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
	// Id of the resource (required)
	Id      *string          `json:"id,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// 16 byte key in hex (32 characters) (required)
	Key *string `json:"key,omitempty"`
	// 16 byte key in hex (32 characters) (required)
	Kid *string `json:"kid,omitempty"`
}

func (m MarlinDrm) DrmType() DrmType {
	return DrmType_MARLIN
}
func (m MarlinDrm) MarshalJSON() ([]byte, error) {
	type M MarlinDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "MARLIN"

	return json.Marshal(x)
}
