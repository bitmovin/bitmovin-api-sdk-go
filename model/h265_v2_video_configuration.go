package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// H265V2VideoConfiguration model
type H265V2VideoConfiguration struct {
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
	PresetConfiguration H265V2PresetConfiguration `json:"presetConfiguration,omitempty"`
	// Automatically configures the encoder for the given SDR/HDR format.
	DynamicRangeFormat H265DynamicRangeFormat `json:"dynamicRangeFormat,omitempty"`
	// Rate control mode configuration. Used to Configure the Perceptual Encoding Mode Settings.
	RateControlModeConfig *H265V2RateControlModeConfig `json:"rateControlModeConfig,omitempty"`
	// Motion compensated temporal filtering mode.
	MotionCompensatedTemporalFiltering H265V2MotionCompensatedTemporalFiltering `json:"motionCompensatedTemporalFiltering,omitempty"`
	// Set codec tier to high when true, main when false.
	LevelHighTier *bool `json:"levelHighTier,omitempty"`
}

func (m H265V2VideoConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_H265V2
}

// UnmarshalJSON unmarshals model H265V2VideoConfiguration from a JSON structure
func (m *H265V2VideoConfiguration) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id                                 *string                                  `json:"id"`
		Name                               *string                                  `json:"name"`
		Description                        *string                                  `json:"description"`
		CreatedAt                          *DateTime                                `json:"createdAt"`
		ModifiedAt                         *DateTime                                `json:"modifiedAt"`
		CustomData                         *map[string]interface{}                  `json:"customData"`
		Width                              *int32                                   `json:"width"`
		Height                             *int32                                   `json:"height"`
		Bitrate                            *int64                                   `json:"bitrate"`
		Rate                               *float64                                 `json:"rate"`
		PixelFormat                        PixelFormat                              `json:"pixelFormat"`
		ColorConfig                        *ColorConfig                             `json:"colorConfig"`
		SampleAspectRatioNumerator         *int32                                   `json:"sampleAspectRatioNumerator"`
		SampleAspectRatioDenominator       *int32                                   `json:"sampleAspectRatioDenominator"`
		DisplayAspectRatio                 *DisplayAspectRatio                      `json:"displayAspectRatio"`
		EncodingMode                       EncodingMode                             `json:"encodingMode"`
		PresetConfiguration                H265V2PresetConfiguration                `json:"presetConfiguration"`
		DynamicRangeFormat                 H265DynamicRangeFormat                   `json:"dynamicRangeFormat"`
		RateControlModeConfig              json.RawMessage                          `json:"rateControlModeConfig"`
		MotionCompensatedTemporalFiltering H265V2MotionCompensatedTemporalFiltering `json:"motionCompensatedTemporalFiltering"`
		LevelHighTier                      *bool                                    `json:"levelHighTier"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result H265V2VideoConfiguration

	result.Id = data.Id
	result.Name = data.Name
	result.Description = data.Description
	result.CreatedAt = data.CreatedAt
	result.ModifiedAt = data.ModifiedAt
	result.CustomData = data.CustomData
	result.Width = data.Width
	result.Height = data.Height
	result.Bitrate = data.Bitrate
	result.Rate = data.Rate
	result.PixelFormat = data.PixelFormat
	result.ColorConfig = data.ColorConfig
	result.SampleAspectRatioNumerator = data.SampleAspectRatioNumerator
	result.SampleAspectRatioDenominator = data.SampleAspectRatioDenominator
	result.DisplayAspectRatio = data.DisplayAspectRatio
	result.EncodingMode = data.EncodingMode
	result.PresetConfiguration = data.PresetConfiguration
	result.DynamicRangeFormat = data.DynamicRangeFormat
	result.MotionCompensatedTemporalFiltering = data.MotionCompensatedTemporalFiltering
	result.LevelHighTier = data.LevelHighTier

	allOfRateControlModeConfig, err := UnmarshalH265V2RateControlModeConfig(bytes.NewBuffer(data.RateControlModeConfig), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.RateControlModeConfig = &allOfRateControlModeConfig

	*m = result

	return nil
}

func (m H265V2VideoConfiguration) MarshalJSON() ([]byte, error) {
	type M H265V2VideoConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "H265V2"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
