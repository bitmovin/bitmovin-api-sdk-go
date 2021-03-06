package model

import (
	"encoding/json"
)

// ChunkedTextMuxing model
type ChunkedTextMuxing struct {
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
	// If this is set and contains objects, then this muxing has been ignored during the encoding process
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Specifies how to handle streams that don't fulfill stream conditions
	StreamConditionsMode StreamConditionsMode `json:"streamConditionsMode,omitempty"`
	// Length of the segments in seconds (required)
	SegmentLength *float64 `json:"segmentLength,omitempty"`
	// Segment naming template
	SegmentNaming *string `json:"segmentNaming,omitempty"`
	// Segment naming template with placeholders which will be replaced during the encoding. The result will be saved in segmentNaming. {rand:4} gets replaced with an alphanumeric string of length specified after the colon. Defaults to 32. If this field is set, segmentNaming must not be specified.
	SegmentNamingTemplate *string `json:"segmentNamingTemplate,omitempty"`
	// Number of segments which have been encoded
	SegmentsMuxed *int32 `json:"segmentsMuxed,omitempty"`
}

func (m ChunkedTextMuxing) MuxingType() MuxingType {
	return MuxingType_CHUNKED_TEXT
}
func (m ChunkedTextMuxing) MarshalJSON() ([]byte, error) {
	type M ChunkedTextMuxing
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CHUNKED_TEXT"

	return json.Marshal(x)
}
