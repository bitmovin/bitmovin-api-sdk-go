package model

// StatisticsPerStream model
type StatisticsPerStream struct {
	// ID of the stream (required)
	StreamId *string `json:"streamId,omitempty"`
	// ID of the condec configuration (required)
	CodecConfigId *string `json:"codecConfigId,omitempty"`
	// Multiplier for the encoded minutes. Depends on muxing type. (required)
	Multiplicator *float64 `json:"multiplicator,omitempty"`
	// Encoded bytes. (required)
	EncodedBytes *int64 `json:"encodedBytes,omitempty"`
	// Length of the stream. (required)
	EncodedSeconds *float64 `json:"encodedSeconds,omitempty"`
	// Minutes you will be charged for (billableMinutes = encodedSeconds * multiplicator) (required)
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	// Video width, only if video stream
	Width *int64 `json:"width,omitempty"`
	// Video height, only if video stream
	Height *int64 `json:"height,omitempty"`
	// If it' a video stream this value is the FPS, for audio it's the sample rate. (required)
	Rate *float64 `json:"rate,omitempty"`
	// Bitrate of the stream (required)
	Bitrate      *int64               `json:"bitrate,omitempty"`
	Codec        CodecConfigType      `json:"codec,omitempty"`
	Resolution   StatisticsResolution `json:"resolution,omitempty"`
	EncodingMode EncodingMode         `json:"encodingMode,omitempty"`
	// The output minutes multiplicator for the given encodingMode
	EncodingModeMultiplicator *float64                 `json:"encodingModeMultiplicator,omitempty"`
	PerTitleResultStream      StatisticsPerTitleStream `json:"perTitleResultStream,omitempty"`
	// The output minutes multiplicator for per-title
	PerTitleMultiplicator *float64          `json:"perTitleMultiplicator,omitempty"`
	PsnrMode              PsnrPerStreamMode `json:"psnrMode,omitempty"`
	// The output minutes multiplicator for psnr streams
	PsnrMultiplicator *float64                 `json:"psnrMultiplicator,omitempty"`
	DolbyVisionMode   DolbyVisionPerStreamMode `json:"dolbyVisionMode,omitempty"`
	// The output minutes multiplicator for Dolby Vision streams
	DolbyVisionMultiplicator *float64 `json:"dolbyVisionMultiplicator,omitempty"`
	// Name of the preset configuration used for the codec configuration or \"CUSTOM\" if any preset values were overridden
	Preset *string `json:"preset,omitempty"`
	// The output minutes multiplicator for the used codec configuration preset.
	PresetMultiplicator *float64 `json:"presetMultiplicator,omitempty"`
	// Indicates if the stream was part of a live encoding.
	Live *bool `json:"live,omitempty"`
	// The output minutes multiplicator for live streams.
	LiveMultiplicator *float64 `json:"liveMultiplicator,omitempty"`
	// Indicates if an enhanced interlace filter was used.
	EnhancedDeinterlace *bool `json:"enhancedDeinterlace,omitempty"`
	// The output minutes multiplicator for streams using an enhanced Deinterlace Filter.
	EnhancedDeinterlaceMultiplicator *float64                       `json:"enhancedDeinterlaceMultiplicator,omitempty"`
	NexGuardABWatermarkingType       *NexGuardAbWatermarkingFeature `json:"nexGuardABWatermarkingType,omitempty"`
	// The output minutes multiplicator for streams using a NexGuard A/B Watermarking.
	NexGuardABWatermarkingMultiplicator *float64            `json:"nexGuardABWatermarkingMultiplicator,omitempty"`
	PixelFormatBitDepth                 PixelFormatBitDepth `json:"pixelFormatBitDepth,omitempty"`
	// The output minutes multiplicator for the pixel format bit depth
	PixelFormatMultiplicator *float64     `json:"pixelFormatMultiplicator,omitempty"`
	InputFactor              *InputFactor `json:"inputFactor,omitempty"`
}
