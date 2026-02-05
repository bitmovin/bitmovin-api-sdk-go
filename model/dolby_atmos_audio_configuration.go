package model

import (
	"encoding/json"
)

// DolbyAtmosAudioConfiguration model
type DolbyAtmosAudioConfiguration struct {
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
	// Target bitrate for the encoded audio in bps. Allowed values are: 384000, 448000, 576000, 640000, 768000, 1024000 (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz. Only 48000 is allowed.
	Rate *float64 `json:"rate,omitempty"`
	// Settings for loudness control (required)
	LoudnessControl *DolbyAtmosLoudnessControl `json:"loudnessControl,omitempty"`
	Preprocessing   *DolbyAtmosPreprocessing   `json:"preprocessing,omitempty"`
}

func (m DolbyAtmosAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DOLBY_ATMOS
}
func (m DolbyAtmosAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DolbyAtmosAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DOLBY_ATMOS"

	return json.Marshal(x)
}
