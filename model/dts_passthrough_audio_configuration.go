package model

import (
	"encoding/json"
)

// DtsPassthroughAudioConfiguration model
type DtsPassthroughAudioConfiguration struct {
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
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
}

func (m DtsPassthroughAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DTS_PASSTHROUGH
}
func (m DtsPassthroughAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DtsPassthroughAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DTS_PASSTHROUGH"

	return json.Marshal(x)
}
