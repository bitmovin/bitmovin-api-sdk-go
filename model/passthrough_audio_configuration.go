package model

import (
	"encoding/json"
)

// PassthroughAudioConfiguration model
type PassthroughAudioConfiguration struct {
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
}

func (m PassthroughAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_AUDIO_PASSTHROUGH
}
func (m PassthroughAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M PassthroughAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AUDIO_PASSTHROUGH"

	return json.Marshal(x)
}
