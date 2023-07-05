package model

import (
	"encoding/json"
)

// H262VideoConfiguration model
type H262VideoConfiguration struct {
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
	// Width of the encoded video in pixels
	Width *int32 `json:"width,omitempty"`
	// Height of the encoded video in pixels
	Height *int32 `json:"height,omitempty"`
	// Target bitrate for the encoded video in bps. Either bitrate or crf is required.
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Target frame rate of the encoded video. Must be set for live encodings
	Rate *float64 `json:"rate,omitempty"`
	// Describes the color encoding, bit depth, and chroma subsampling of each pixel in the output image.
	PixelFormat PixelFormat  `json:"pixelFormat,omitempty"`
	ColorConfig *ColorConfig `json:"colorConfig,omitempty"`
	// The numerator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioDenominator is set. If set then displayAspectRatio is not allowed.
	SampleAspectRatioNumerator *int32 `json:"sampleAspectRatioNumerator,omitempty"`
	// The denominator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioNumerator is set. If set then displayAspectRatio is not allowed.
	SampleAspectRatioDenominator *int32 `json:"sampleAspectRatioDenominator,omitempty"`
	// Specifies a display aspect ratio (DAR) to be enforced. The sample aspect ratio (SAR) will be adjusted accordingly. If set then sampleAspectRatioNumerator and sampleAspectRatioDenominator are not allowed.
	DisplayAspectRatio *DisplayAspectRatio `json:"displayAspectRatio,omitempty"`
	// The mode of the encoding. When this is set, `encodingMode` (`liveEncodingMode`) must not be set in the (live) encoding start request.
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	// Use a set of well defined configurations preset to support certain use cases. Can be overwritten with more specific values.
	PresetConfiguration H262PresetConfiguration `json:"presetConfiguration,omitempty"`
	// When setting a profile, all other settings must not exceed the limits which are defined in the profile. Otherwise, a higher profile may be automatically chosen. (required)
	Profile ProfileH262 `json:"profile,omitempty"`
	// Amount of b frames.
	Bframes *int32 `json:"bframes,omitempty"`
	// Maximum Bitrate
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// Minimum Bitrate
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Playback device buffer size
	Bufsize *int64 `json:"bufsize,omitempty"`
	// Minimum GOP length, the minimum distance between I-frames
	GopSize *int32 `json:"gopSize,omitempty"`
	// Specified set of constraints that indicate a degree of required decoder performance for a profile
	Level LevelH262 `json:"level,omitempty"`
	// Using TOP_FIELD_FIRST or BOTTOM_FIELD_FIRST will output interlaced video
	InterlaceMode H262InterlaceMode `json:"interlaceMode,omitempty"`
}

func (m H262VideoConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_H262
}
func (m H262VideoConfiguration) MarshalJSON() ([]byte, error) {
	type M H262VideoConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "H262"

	return json.Marshal(x)
}
