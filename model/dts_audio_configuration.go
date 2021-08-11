package model

import (
	"encoding/json"
)

// DtsAudioConfiguration model
type DtsAudioConfiguration struct {
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
	// Target bitrate for the encoded audio in bps. Allowed values are: 255000, 384000, for DTS Express Audio. For DTS Digital Surround only 768000 is allowed. (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz. Only 48000 is allowed.
	Rate *float64 `json:"rate,omitempty"`
	// There are two DTS modes available: DTS Express Audio (EXPRESS_AUDIO) and DTS Digital Surround (DIGITAL_SURROUND)
	Mode DtsMode `json:"mode,omitempty"`
}

func (m DtsAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DTS
}
func (m DtsAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DtsAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DTS"

	return json.Marshal(x)
}
