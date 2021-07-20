package model

import (
	"encoding/json"
)

// PcmAudioConfiguration model
type PcmAudioConfiguration struct {
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
	// Target bitrate for the encoded audio in bps (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz
	Rate *float64 `json:"rate,omitempty"`
	// Use a set of well defined configurations preset to support certain use cases. Can be overwritten with more specific values.
	PresetConfiguration PcmPresetConfiguration `json:"presetConfiguration,omitempty"`
	// Channel layout of the audio codec configuration
	ChannelLayout PcmChannelLayout `json:"channelLayout,omitempty"`
	// Sampling format of the audio codec configuration
	SampleFormat PcmSampleFormat `json:"sampleFormat,omitempty"`
}

func (m PcmAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_PCM
}
func (m PcmAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M PcmAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "PCM"

	return json.Marshal(x)
}
