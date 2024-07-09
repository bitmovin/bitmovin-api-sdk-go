package model

import (
	"encoding/json"
)

// H264VideoConfiguration model
type H264VideoConfiguration struct {
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
	// Choose from a set of preset configurations tailored for common use cases. Check out [H264 Presets](https://bitmovin.com/docs/encoding/tutorials/h264-presets) to see which values get applied by each preset. Explicitly setting a property to a different value will override the preset's value for that property.
	PresetConfiguration PresetConfiguration `json:"presetConfiguration,omitempty"`
	// Automatically configures the H264 Video Codec to be compatible with the given SDR format. Bitmovin recommends to use the dynamic range format together with a preset configuration to achieve good results. Explicitly configured properties will take precedence over dynamic range format settings, which in turn will take precedence over preset configurations.
	DynamicRangeFormat H264DynamicRangeFormat `json:"dynamicRangeFormat,omitempty"`
	// Constant rate factor for quality-based variable bitrate. Either bitrate or crf is required.
	Crf *float64 `json:"crf,omitempty"`
	// When setting a profile, all other settings must not exceed the limits which are defined in the profile. Otherwise, a higher profile may be automatically chosen. (required)
	Profile ProfileH264 `json:"profile,omitempty"`
	// Amount of b frames
	Bframes *int32 `json:"bframes,omitempty"`
	// Amount of reference frames.
	RefFrames *int32 `json:"refFrames,omitempty"`
	// Minimum quantization factor
	QpMin *int32 `json:"qpMin,omitempty"`
	// Maximum quantization factor
	QpMax            *int32           `json:"qpMax,omitempty"`
	MvPredictionMode MvPredictionMode `json:"mvPredictionMode,omitempty"`
	// Maximum motion vector search range
	MvSearchRangeMax *int32 `json:"mvSearchRangeMax,omitempty"`
	// Enable or disable CABAC
	Cabac *bool `json:"cabac,omitempty"`
	// Maximum Bitrate (bps)
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// Minimum Bitrate (bps)
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Playback device buffer size (bits)
	Bufsize *int64 `json:"bufsize,omitempty"`
	// Minimum GOP length, the minimum distance between I-frames
	MinGop *int32 `json:"minGop,omitempty"`
	// Maximum GOP length, the maximum distance between I-frames
	MaxGop *int32 `json:"maxGop,omitempty"`
	// Enable open-gop, allows referencing frames from a previous gop
	OpenGop *bool `json:"openGop,omitempty"`
	// Minimum interval in seconds between key frames
	MinKeyframeInterval *float64 `json:"minKeyframeInterval,omitempty"`
	// Maximum interval in seconds between key frames
	MaxKeyframeInterval *float64 `json:"maxKeyframeInterval,omitempty"`
	// If three-pass encoding is used and a level is set for the encoder, the bitrate for some segments may exceed the bitrate limit which is defined by the level.
	Level                  LevelH264                  `json:"level,omitempty"`
	BAdaptiveStrategy      BAdapt                     `json:"bAdaptiveStrategy,omitempty"`
	MotionEstimationMethod H264MotionEstimationMethod `json:"motionEstimationMethod,omitempty"`
	// Number of frames for frame-type decision lookahead
	RcLookahead *int32 `json:"rcLookahead,omitempty"`
	// Subpixel motion estimation and mode decision
	SubMe H264SubMe `json:"subMe,omitempty"`
	// Enables or disables Trellis quantization. NOTE: This requires cabac
	Trellis H264Trellis `json:"trellis,omitempty"`
	// Partitions to consider. Analyzing more partition options improves quality at the cost of speed.
	Partitions []H264Partition `json:"partitions,omitempty"`
	// Number of slices per frame.
	Slices *int32 `json:"slices,omitempty"`
	// Using TOP_FIELD_FIRST or BOTTOM_FIELD_FIRST will output interlaced video
	InterlaceMode H264InterlaceMode `json:"interlaceMode,omitempty"`
	// Scene change sensitivity. The higher the value, the more likely an I-frame will be inserted. Set to 0 to disable it which is advised for scenarios where fixed GOP is required, e.g., adaptive streaming outputs like DASH, HLS and Smooth. Having this setting enabled can improve quality for progressive output with an increased internal chunk length (see `internalChunkLength` of muxings).
	SceneCutThreshold *int32 `json:"sceneCutThreshold,omitempty"`
	// Signal hypothetical reference decoder (HRD) information (requires bufsize to be set)
	NalHrd H264NalHrd `json:"nalHrd,omitempty"`
	// Keep some B-frames as references
	BPyramid H264BPyramid `json:"bPyramid,omitempty"`
	// Defines whether CEA 608/708 subtitles are copied from the input video stream
	Cea608708SubtitleConfig *Cea608708SubtitleConfiguration `json:"cea608708SubtitleConfig,omitempty"`
	// Strength of the in-loop deblocking filter. Higher values deblock more effectively but also soften the image
	DeblockAlpha *int32 `json:"deblockAlpha,omitempty"`
	// Threshold of the in-loop deblocking filter. Higher values apply deblocking stronger on non flat blocks, lower values on flat blocks
	DeblockBeta *int32 `json:"deblockBeta,omitempty"`
	// Controls the adaptive quantization algorithm
	AdaptiveQuantizationMode AdaptiveQuantMode `json:"adaptiveQuantizationMode,omitempty"`
	// Values greater than 1 reduce blocking and blurring in flat and textured areas. Values less than 1 reduces ringing artifacts at the cost of more banding artifacts. Negative values are not allowed
	AdaptiveQuantizationStrength *float64 `json:"adaptiveQuantizationStrength,omitempty"`
	// Allow references on a per partition basis, rather than per-macroblock basis
	MixedReferences *bool `json:"mixedReferences,omitempty"`
	// Enables adaptive spatial transform (high profile 8x8 transform)
	AdaptiveSpatialTransform *bool `json:"adaptiveSpatialTransform,omitempty"`
	// Enables fast skip detection on P-frames. Disabling this very slightly increases quality but at a large speed loss
	FastSkipDetectionPFrames *bool `json:"fastSkipDetectionPFrames,omitempty"`
	// Enable open-gop, allows referencing frames from a previous gop
	WeightedPredictionBFrames *bool `json:"weightedPredictionBFrames,omitempty"`
	// Defines the mode for weighted prediction for P-frames
	WeightedPredictionPFrames WeightedPredictionPFrames `json:"weightedPredictionPFrames,omitempty"`
	// Enable macroblock tree ratecontrol. Macroblock tree rate control tracks how often blocks of the frame are used for prediciting future frames
	MacroblockTreeRatecontrol *bool `json:"macroblockTreeRatecontrol,omitempty"`
	// Ratio between constant bitrate (0.0) and constant quantizer (1.0). Valid range 0.0 - 1.0
	QuantizerCurveCompression *float64 `json:"quantizerCurveCompression,omitempty"`
	// Psychovisual Rate Distortion retains fine details like film grain at the expense of more blocking artifacts. Higher values make the video appear sharper and more detailed but with a higher risk of blocking artifacts. Needs to have subMe with RD_IP, RD_ALL, RD_REF_IP or RD_REF_ALL
	PsyRateDistortionOptimization *float64 `json:"psyRateDistortionOptimization,omitempty"`
	// Higher values will improve sharpness and detail retention but might come at costs of artifacts. Needs to have trellis enabled
	PsyTrellis *float64 `json:"psyTrellis,omitempty"`
	// Enable/disable automatic calculation of level, maxBitrate, and bufsize based on the least level that satisfies maximum property values for picture resolution, frame rate, and bit rate. In the case the target level is set explicitly, the maximum bitrate and buffer size are calculated based on the defined level. Explicitly setting maxBitrate, or bufsize properties will disable the automatic calculation.
	AutoLevelSetup AutoLevelSetup `json:"autoLevelSetup,omitempty"`
}

func (m H264VideoConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_H264
}
func (m H264VideoConfiguration) MarshalJSON() ([]byte, error) {
	type M H264VideoConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "H264"

	return json.Marshal(x)
}
