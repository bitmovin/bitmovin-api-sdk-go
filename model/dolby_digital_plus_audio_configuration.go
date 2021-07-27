package model

import (
	"encoding/json"
)

// DolbyDigitalPlusAudioConfiguration model
type DolbyDigitalPlusAudioConfiguration struct {
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
	// BitstreamInfo defines metadata parameters contained in the Dolby Digital Plus audio bitstream
	BitstreamInfo *DolbyDigitalPlusBitstreamInfo `json:"bitstreamInfo,omitempty"`
	// Channel layout of the audio codec configuration.
	ChannelLayout DolbyDigitalPlusChannelLayout `json:"channelLayout,omitempty"`
	Downmixing    *DolbyDigitalPlusDownmixing   `json:"downmixing,omitempty"`
	// It provides a framework for signaling new evolution framework applications, such as Intelligent Loudness, in each Dolby codec.
	EvolutionFrameworkControl DolbyDigitalPlusEvolutionFrameworkControl `json:"evolutionFrameworkControl,omitempty"`
	// Settings for loudness control (required)
	LoudnessControl *DolbyDigitalPlusLoudnessControl `json:"loudnessControl,omitempty"`
	Preprocessing   *DolbyDigitalPlusPreprocessing   `json:"preprocessing,omitempty"`
}

func (m DolbyDigitalPlusAudioConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DDPLUS
}
func (m DolbyDigitalPlusAudioConfiguration) MarshalJSON() ([]byte, error) {
	type M DolbyDigitalPlusAudioConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DDPLUS"

	return json.Marshal(x)
}
