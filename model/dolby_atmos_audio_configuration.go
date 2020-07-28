package model
import (
	"time"
)

type DolbyAtmosAudioConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Target bitrate for the encoded audio in bps. Allowed values are: 384000, 448000, 576000, 640000, 768000, 1024000 (required)
	Bitrate *map[string]interface{} `json:"bitrate,omitempty"`
	// Audio sampling rate in Hz. Only 48000 is allowed.
	Rate *map[string]interface{} `json:"rate,omitempty"`
	// Settings for loudness control (required)
	LoudnessControl *DolbyAtmosLoudnessControl `json:"loudnessControl,omitempty"`
}
func (o DolbyAtmosAudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_DOLBY_ATMOS
}

