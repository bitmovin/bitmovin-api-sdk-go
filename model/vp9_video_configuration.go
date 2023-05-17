package model

import (
	"encoding/json"
)

// Vp9VideoConfiguration model
type Vp9VideoConfiguration struct {
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
	// The mode of the encoding
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	// Choose from a set of preset configurations tailored for common use cases. Check out [VP9 Presets](https://bitmovin.com/docs/encoding/tutorials/vp9-presets) to see which values get applied by each preset. Explicitly setting a property to a different value will override the preset's value for that property.
	PresetConfiguration PresetConfiguration `json:"presetConfiguration,omitempty"`
	// Automatically configures the VP9 Video Codec to be compatible with the given SDR/HLG format. Bitmovin recommends to use the dynamic range format together with a preset configuration to achieve good results. Explicitly configured properties will take precedence over dynamic range format settings, which in turn will take precedence over preset configurations.
	DynamicRangeFormat Vp9DynamicRangeFormat `json:"dynamicRangeFormat,omitempty"`
	// Constant rate factor for quality-based variable bitrate. Either bitrate or crf is required.
	Crf *int32 `json:"crf,omitempty"`
	// Number of frames to look ahead for alternate reference frame selection.
	LagInFrames *int32 `json:"lagInFrames,omitempty"`
	// Enables error resiliency feature
	ErrorResiliencyEnabled *bool `json:"errorResiliencyEnabled,omitempty"`
	// Number of tile columns to use, log2. Depending on the encoding width there are limitations on this value. The minimum values are 2 for width >= 1920 and 1 for width >= 1280. The minimum width of each tile is 256 pixels so the maximum values are 0 for width < 256, 1 for width < 512, 2 for width < 1024, 3 for width < 2048, 4 for width < 4096, 5 for width < 8192. If the value is too high or too low it will be overridden.
	TileColumns *int32 `json:"tileColumns,omitempty"`
	// Number of tile rows to use, log2.
	TileRows *int32 `json:"tileRows,omitempty"`
	// Enable frame parallel decodability features
	FrameParallel *bool `json:"frameParallel,omitempty"`
	// Maximum I-frame bitrate (percentage) 0=unlimited
	MaxIntraRate *int64 `json:"maxIntraRate,omitempty"`
	// Minimum quantization factor.
	QpMin *int32 `json:"qpMin,omitempty"`
	// Maximum quantization factor.
	QpMax *int32 `json:"qpMax,omitempty"`
	// Datarate undershoot (min) target (percentage).
	RateUndershootPct *int32 `json:"rateUndershootPct,omitempty"`
	// Datarate overshoot (max) target (percentage).
	RateOvershootPct *int32 `json:"rateOvershootPct,omitempty"`
	// Client buffer size (ms)
	ClientBufferSize *int64 `json:"clientBufferSize,omitempty"`
	// Client initial buffer size (ms)
	ClientInitialBufferSize *int64 `json:"clientInitialBufferSize,omitempty"`
	// CBR/VBR bias (0=CBR, 100=VBR)
	BiasPct *int32 `json:"biasPct,omitempty"`
	// Enable noise sensitivity on Y channel
	NoiseSensitivity *bool `json:"noiseSensitivity,omitempty"`
	// Controls the tradeoff between compression efficiency and encoding speed. Higher values indicate a faster encoding. The minimum value for width * height >= 1280 * 720 is 2. If the value is too low it will be overridden.
	CpuUsed *int32 `json:"cpuUsed,omitempty"`
	// Enable automatic alternate reference frames (2pass only)
	AutomaticAltRefFramesEnabled *bool `json:"automaticAltRefFramesEnabled,omitempty"`
	// Target level (255: off, 0: only keep level stats; 10: level 1.0; 11: level 1.1; ... 62: level 6.2)
	TargetLevel *int32 `json:"targetLevel,omitempty"`
	// Enable row based non-deterministic multi-threading
	RowMultiThreadingEnabled *bool `json:"rowMultiThreadingEnabled,omitempty"`
	// Loop filter sharpness.
	Sharpness *int32 `json:"sharpness,omitempty"`
	// Minimum GOP length, the minimum distance between I-frames.
	MinGop *int32 `json:"minGop,omitempty"`
	// Maximum GOP length, the maximum distance between I-frames
	MaxGop *int32 `json:"maxGop,omitempty"`
	// Minimum interval in seconds between key frames
	MinKeyframeInterval *float64 `json:"minKeyframeInterval,omitempty"`
	// Maximum interval in seconds between key frames
	MaxKeyframeInterval *float64   `json:"maxKeyframeInterval,omitempty"`
	Quality             Vp9Quality `json:"quality,omitempty"`
	// Lossless mode
	Lossless *bool `json:"lossless,omitempty"`
	// A change threshold on blocks below which they will be skipped by the encoder.
	StaticThresh *int32    `json:"staticThresh,omitempty"`
	AqMode       Vp9AqMode `json:"aqMode,omitempty"`
	// altref noise reduction max frame count.
	ArnrMaxFrames *int32 `json:"arnrMaxFrames,omitempty"`
	// altref noise reduction filter strength.
	ArnrStrength *int32      `json:"arnrStrength,omitempty"`
	ArnrType     Vp9ArnrType `json:"arnrType,omitempty"`
	// Enable/disable automatic calculation of level, maxBitrate, and bufsize based on the least level that satisfies maximum property values for picture resolution, frame rate, and bit rate. In the case the target level is set explicitly, the maximum bitrate and buffer size are calculated based on the defined level. Explicitly setting rateOvershootPct, or clientBufferSize properties will disable the automatic calculation.
	AutoLevelSetup AutoLevelSetup `json:"autoLevelSetup,omitempty"`
}

func (m Vp9VideoConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_VP9
}
func (m Vp9VideoConfiguration) MarshalJSON() ([]byte, error) {
	type M Vp9VideoConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "VP9"

	return json.Marshal(x)
}
