package model
import (
	"time"
)

type PcmAudioConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
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
func (o PcmAudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_PCM
}

