package model

import (
	"encoding/json"
)

// RtmpInput model
type RtmpInput struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Is set automatically
	Name *string `json:"name,omitempty"`
	// Description for the resource. Is set automatically
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
}

func (m RtmpInput) InputType() InputType {
	return InputType_RTMP
}
func (m RtmpInput) MarshalJSON() ([]byte, error) {
	type M RtmpInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "RTMP"

	return json.Marshal(x)
}
