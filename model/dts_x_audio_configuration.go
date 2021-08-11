package model

import (
	"encoding/json"
)

// The configuration for the DTS:X
type DtsXAudioConfiguration struct {
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
	Bitrate    MediaConfigBitrate      `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz. Must be 48000
	Rate          *float64          `json:"rate,omitempty"`
	ChannelLayout DtsXChannelLayout `json:"channelLayout,omitempty"`
	// Loudness relative to full scale (K-weighted).
	LkfsValue       *float64        `json:"lkfsValue,omitempty"`
	OttLoudnessMode OttLoudnessMode `json:"ottLoudnessMode,omitempty"`
	// Specifies the sync interval which ranges from 1 to 10
	SyncInterval *int64 `json:"syncInterval,omitempty"`
}

func (m DtsXAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DTSX
}
func (m DtsXAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DtsXAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DTSX"

	return json.Marshal(x)
}
