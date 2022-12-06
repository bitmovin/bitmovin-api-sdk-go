package model

import (
	"encoding/json"
)

// HeAacV1AudioConfiguration model
type HeAacV1AudioConfiguration struct {
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
	// Spectral Band Replication (SBR) and Parameteric Stereo (PS) signaling style.
	Signaling HeAacV1Signaling `json:"signaling,omitempty"`
}

func (m HeAacV1AudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_HE_AAC_V1
}
func (m HeAacV1AudioConfiguration) MarshalJSON() ([]byte, error) {
	type M HeAacV1AudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "HE_AAC_V1"

	return json.Marshal(x)
}
