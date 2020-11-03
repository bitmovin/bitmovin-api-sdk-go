package model

import (
    "encoding/json"
)

// Fmp4Muxing model
type Fmp4Muxing struct {
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
    // Length of the fragments in seconds (required)
    SegmentLength *float64 `json:"segmentLength,omitempty"`
    // Segment naming policy
    SegmentNaming *string `json:"segmentNaming,omitempty"`
    // Segment naming policy containing a placeholder of the format '{rand_chars:x}', which will be replaced by a random alphanumeric string of length x (default 32) on each (re)start of the encoding. The resulting string will be copied to the segmentNaming property. Intended to avoid re-use of segment names after restarting a live encoding. If segmentNamingTemplate is set, segmentNaming must not be set.
    SegmentNamingTemplate *string `json:"segmentNamingTemplate,omitempty"`
    // Init segment name
    InitSegmentName *string `json:"initSegmentName,omitempty"`
    // Segment naming policy containing a placeholder of the format '{rand_chars:x}', which will be replaced by a random alphanumeric string of length x (default 32) on each (re)start of the encoding. The resulting string will be copied to the initSegmentName property. Intended to avoid re-use of segment names after restarting a live encoding. If initSegmentNameTemplate is set, initSegmentName must not be set.
    InitSegmentNameTemplate *string `json:"initSegmentNameTemplate,omitempty"`
    // Writes the duration per sample into the sample entry in the Track Fragment Run Box. This could help to fix playback issues on legacy players. Enabling this flag increases the muxing overhead by 4 bytes per sample/frame.
    WriteDurationPerSample *bool `json:"writeDurationPerSample,omitempty"`
    // Number of segments which have been encoded
    SegmentsMuxed *int32 `json:"segmentsMuxed,omitempty"`
}

func (m Fmp4Muxing) MuxingType() MuxingType {
    return MuxingType_FMP4
}
func (m Fmp4Muxing) MarshalJSON() ([]byte, error) {
    type M Fmp4Muxing
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "FMP4"

    return json.Marshal(x)
}


