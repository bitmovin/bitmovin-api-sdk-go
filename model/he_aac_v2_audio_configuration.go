package model

import (
	"encoding/json"
)

// HeAacV2AudioConfiguration model
type HeAacV2AudioConfiguration struct {
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
	// Channel layout of the audio codec configuration
	ChannelLayout AacChannelLayout `json:"channelLayout,omitempty"`
	// Sets the Spectral Band Replication (SBR) and Parameteric Stereo (PS) signaling style.
	Signaling HeAacV2Signaling `json:"signaling,omitempty"`
}

func (m HeAacV2AudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_HE_AAC_V2
}
func (m HeAacV2AudioConfiguration) MarshalJSON() ([]byte, error) {
	type M HeAacV2AudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "HE_AAC_V2"

	return json.Marshal(x)
}
