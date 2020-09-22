package model

import (
	"encoding/json"
)

// PlayReadyDrm model
type PlayReadyDrm struct {
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
	// 16 byte encryption key, 32 hexadecimal characters. Either key or keySeed is required
	Key *string `json:"key,omitempty"`
	// Key seed to generate key. Either key or keySeed is required
	KeySeed *string `json:"keySeed,omitempty"`
	// URL of the license server
	LaUrl *string `json:"laUrl,omitempty"`
	// Base64 encoded pssh payload
	Pssh   *string                   `json:"pssh,omitempty"`
	Method PlayReadyEncryptionMethod `json:"method,omitempty"`
	// Key identifier
	Kid                   *string                         `json:"kid,omitempty"`
	AdditionalInformation *PlayReadyAdditionalInformation `json:"additionalInformation,omitempty"`
}

func (m PlayReadyDrm) DrmType() DrmType {
	return DrmType_PLAYREADY
}
func (m PlayReadyDrm) MarshalJSON() ([]byte, error) {
	type M PlayReadyDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "PLAYREADY"

	return json.Marshal(x)
}
