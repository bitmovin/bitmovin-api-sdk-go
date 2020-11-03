package model

import (
    "encoding/json"
)

// OpusAudioConfiguration model
type OpusAudioConfiguration struct {
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
    ChannelLayout OpusChannelLayout `json:"channelLayout,omitempty"`
}

func (m OpusAudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_OPUS
}
func (m OpusAudioConfiguration) MarshalJSON() ([]byte, error) {
    type M OpusAudioConfiguration
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "OPUS"

    return json.Marshal(x)
}


