package model

import (
	"encoding/json"
)

// AesEncryptionDrm model
type AesEncryptionDrm struct {
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
	// 16 byte initialization vector
	Iv *string `json:"iv,omitempty"`
	// Path relative to the output for referencing in the manifest. If this value is not set the key file will be written automatically to the output folder.
	KeyFileUri *string             `json:"keyFileUri,omitempty"`
	Method     AesEncryptionMethod `json:"method,omitempty"`
}

func (m AesEncryptionDrm) DrmType() DrmType {
	return DrmType_AES
}
func (m AesEncryptionDrm) MarshalJSON() ([]byte, error) {
	type M AesEncryptionDrm
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AES"

	return json.Marshal(x)
}
