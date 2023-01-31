package model

import (
	"encoding/json"
)

// ProgressiveTsMuxing model
type ProgressiveTsMuxing struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
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
	Streams    []MuxingStream          `json:"streams,omitempty"`
	Outputs    []EncodingOutput        `json:"outputs,omitempty"`
	// Average bitrate. Available after encoding finishes.
	AvgBitrate *int64 `json:"avgBitrate,omitempty"`
	// Min bitrate. Available after encoding finishes.
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Max bitrate. Available after encoding finishes.
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// This read-only property is set during the analysis step of the encoding. If it contains items, the Muxing has been ignored during the encoding process according to its 'streamConditionsMode'
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Specifies how to proceed with the Muxing when some of its Streams are ignored (see 'condition' property of the Stream resource). The settings only make a difference for Muxings with more than one Stream. When retrieving the resource after the analysis step of the encoding has finished, 'ignoredBy' will indicate if and why it has been ignored.
	StreamConditionsMode StreamConditionsMode `json:"streamConditionsMode,omitempty"`
	// Length of the segments in seconds
	SegmentLength *float64 `json:"segmentLength,omitempty"`
	// Name of the output file
	Filename *string `json:"filename,omitempty"`
	// Offset of MPEG-TS timestamps in seconds. e.g. first packet will start with PTS 900,000 for a 10 seconds offset (90,000 MPEG-TS timescale).
	StartOffset *int32 `json:"startOffset,omitempty"`
}

func (m ProgressiveTsMuxing) MuxingType() MuxingType {
	return MuxingType_PROGRESSIVE_TS
}
func (m ProgressiveTsMuxing) MarshalJSON() ([]byte, error) {
	type M ProgressiveTsMuxing
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "PROGRESSIVE_TS"

	return json.Marshal(x)
}
