package model

import (
	"encoding/json"
)

// AudioMixInputStream model
type AudioMixInputStream struct {
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
	// Channel layout of the audio mix input stream
	ChannelLayout    AudioMixInputChannelLayout   `json:"channelLayout,omitempty"`
	AudioMixChannels []AudioMixInputStreamChannel `json:"audioMixChannels,omitempty"`
}

func (m AudioMixInputStream) InputStreamType() InputStreamType {
	return InputStreamType_AUDIO_MIX
}
func (m AudioMixInputStream) MarshalJSON() ([]byte, error) {
	type M AudioMixInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AUDIO_MIX"

	return json.Marshal(x)
}
