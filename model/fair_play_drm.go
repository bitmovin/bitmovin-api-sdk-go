package model

import (
	"encoding/json"
)

// FairPlayDrm model
type FairPlayDrm struct {
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
	// 16 byte Encryption key, 32 hexadecimal characters (required)
	Key *string `json:"key,omitempty"`
	// 16 byte initialization vector (required)
	Iv *string `json:"iv,omitempty"`
	// Url of the licensing server
	Uri *string `json:"uri,omitempty"`
}

func (m FairPlayDrm) DrmType() DrmType {
	return DrmType_FAIRPLAY
}
func (m FairPlayDrm) MarshalJSON() ([]byte, error) {
	type M FairPlayDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "FAIRPLAY"

	return json.Marshal(x)
}
