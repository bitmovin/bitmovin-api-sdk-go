package model

import (
	"encoding/json"
)

// WidevineDrm model
type WidevineDrm struct {
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
	Outputs    []EncodingOutput        `json:"outputs,omitempty"`
	// 16 byte Encryption key, 32 hexadecimal characters (required)
	Key *string `json:"key,omitempty"`
	// 16 byte Key id, 32 hexadecimal characters (required)
	Kid *string `json:"kid,omitempty"`
	// Base 64 Encoded (required)
	Pssh *string `json:"pssh,omitempty"`
}

func (m WidevineDrm) DrmType() DrmType {
	return DrmType_WIDEVINE
}
func (m WidevineDrm) MarshalJSON() ([]byte, error) {
	type M WidevineDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "WIDEVINE"

	return json.Marshal(x)
}
