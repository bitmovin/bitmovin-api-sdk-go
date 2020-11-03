package model

import (
    "encoding/json"
)

// AudioMixFilter model
type AudioMixFilter struct {
    // Name of the resource. Can be freely chosen by the user.
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
    // Channel layout of the audio codec configuration (required)
    ChannelLayout AudioMixChannelLayout `json:"channelLayout,omitempty"`
    // List of mixed channels that matches the channel layout (required)
    AudioMixChannels []AudioMixChannel `json:"audioMixChannels,omitempty"`
}

func (m AudioMixFilter) FilterType() FilterType {
    return FilterType_AUDIO_MIX
}
func (m AudioMixFilter) MarshalJSON() ([]byte, error) {
    type M AudioMixFilter
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "AUDIO_MIX"

    return json.Marshal(x)
}


