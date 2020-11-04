package model

import (
	"encoding/json"
)

// ClearKeyDrm model
type ClearKeyDrm struct {
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
	// 16 byte encryption key, 32 hexadecimal characters (required)
	Key *string `json:"key,omitempty"`
	// 16 byte key id (required)
	Kid *string `json:"kid,omitempty"`
}

func (m ClearKeyDrm) DrmType() DrmType {
	return DrmType_CLEARKEY
}
func (m ClearKeyDrm) MarshalJSON() ([]byte, error) {
	type M ClearKeyDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CLEARKEY"

	return json.Marshal(x)
}
