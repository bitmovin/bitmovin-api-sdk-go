package model

import (
    "encoding/json"
)

// AacAudioConfiguration model
type AacAudioConfiguration struct {
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
    // Id of the resource (required)
    Id *string `json:"id,omitempty"`
    // Target bitrate for the encoded audio in bps (required)
    Bitrate *int64 `json:"bitrate,omitempty"`
    // Audio sampling rate in Hz
    Rate *float64 `json:"rate,omitempty"`
    // Channel layout of the audio codec configuration
    ChannelLayout AacChannelLayout `json:"channelLayout,omitempty"`
    // The highest frequency that will pass the audio encoder. This value is optional.
    CutoffFrequency *int32 `json:"cutoffFrequency,omitempty"`
}

func (m AacAudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_AAC
}
func (m AacAudioConfiguration) MarshalJSON() ([]byte, error) {
    type M AacAudioConfiguration
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "AAC"

    return json.Marshal(x)
}


