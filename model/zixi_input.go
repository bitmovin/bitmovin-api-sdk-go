package model

import (
	"encoding/json"
)

// ZixiInput model
type ZixiInput struct {
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
	CustomData     *map[string]interface{} `json:"customData,omitempty"`
	Host           *string                 `json:"host,omitempty"`
	Port           *int32                  `json:"port,omitempty"`
	Stream         *string                 `json:"stream,omitempty"`
	Password       *string                 `json:"password,omitempty"`
	Latency        *int32                  `json:"latency,omitempty"`
	MinBitrate     *int32                  `json:"minBitrate,omitempty"`
	DecryptionType *string                 `json:"decryptionType,omitempty"`
	DecryptionKey  *string                 `json:"decryptionKey,omitempty"`
}

func (m ZixiInput) InputType() InputType {
	return InputType_ZIXI
}
func (m ZixiInput) MarshalJSON() ([]byte, error) {
	type M ZixiInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "ZIXI"

	return json.Marshal(x)
}
