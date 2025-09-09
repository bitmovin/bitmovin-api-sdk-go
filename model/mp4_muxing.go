package model

import (
	"encoding/json"
)

// Mp4Muxing model
type Mp4Muxing struct {
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
	// Name of the output file (either `filename` or `name` is required, prefer `filename`)
	Filename *string `json:"filename,omitempty"`
	//  Duration of fragments in milliseconds. Required for Fragmented MP4 Muxing (for Smooth Streaming or DASH On-Demand). Not setting this will result in unfragmented mp4.
	FragmentDuration *int32 `json:"fragmentDuration,omitempty"`
	// Prevents creation of very short fragments (in seconds). If the last fragment is shorter than minimumFragmentDuration or there is a custom keyframe too close to a fragment boundary, short fragments will be omitted by removing fragment boundaries, resulting in a fragment of a size up to fragmentDuration + minimumFragmentDuration.
	MinimumFragmentDuration         *float64                        `json:"minimumFragmentDuration,omitempty"`
	FragmentedMP4MuxingManifestType FragmentedMp4MuxingManifestType `json:"fragmentedMP4MuxingManifestType,omitempty"`
	// Dolby Vision specific configuration
	DolbyVisionConfiguration *DolbyVisionMuxingConfiguration `json:"dolbyVisionConfiguration,omitempty"`
}

func (m Mp4Muxing) MuxingType() MuxingType {
	return MuxingType_MP4
}
func (m Mp4Muxing) MarshalJSON() ([]byte, error) {
	type M Mp4Muxing
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "MP4"

	return json.Marshal(x)
}
