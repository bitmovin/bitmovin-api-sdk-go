package model

import (
    "encoding/json"
)

// Mp3Muxing model
type Mp3Muxing struct {
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
    Streams []MuxingStream `json:"streams,omitempty"`
    Outputs []EncodingOutput `json:"outputs,omitempty"`
    // Average bitrate. Available after encoding finishes.
    AvgBitrate *int64 `json:"avgBitrate,omitempty"`
    // Min bitrate. Available after encoding finishes.
    MinBitrate *int64 `json:"minBitrate,omitempty"`
    // Max bitrate. Available after encoding finishes.
    MaxBitrate *int64 `json:"maxBitrate,omitempty"`
    // If this is set and contains objects, then this muxing has been ignored during the encoding process
    IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
    // Specifies how to handle streams that don't fulfill stream conditions
    StreamConditionsMode StreamConditionsMode `json:"streamConditionsMode,omitempty"`
    // Name of the new file (required)
    Filename *string `json:"filename,omitempty"`
}

func (m Mp3Muxing) MuxingType() MuxingType {
    return MuxingType_MP3
}
func (m Mp3Muxing) MarshalJSON() ([]byte, error) {
    type M Mp3Muxing
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "MP3"

    return json.Marshal(x)
}


