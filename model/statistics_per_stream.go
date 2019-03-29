package model

type StatisticsPerStream struct {
	// ID of the stream
	StreamId string `json:"streamId,omitempty"`
	// ID of the condec configuration
	CodecConfigId string `json:"codecConfigId,omitempty"`
	// Multiplier for the encoded minutes. Depends on muxing type.
	Multiplicator *float64 `json:"multiplicator,omitempty"`
	// Encoded bytes.
	EncodedBytes *int64 `json:"encodedBytes,omitempty"`
	// Length of the stream.
	EncodedSeconds *float64 `json:"encodedSeconds,omitempty"`
	// Minutes you will be charged for (billableMinutes = encodedSeconds * multiplicator)
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	// Video width, only if video stream
	Width *int64 `json:"width,omitempty"`
	// Video height, only if video stream
	Height *int64 `json:"height,omitempty"`
	// If it' a video stream this value is the FPS, for audio it's the sample rate.
	Rate *int64 `json:"rate,omitempty"`
	// Bitrate of the stream
	Bitrate *int64 `json:"bitrate,omitempty"`
	Codec CodecConfigType `json:"codec,omitempty"`
	Resolution StatisticsResolution `json:"resolution,omitempty"`
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	// The output minutes multiplicator for the given encodingMode
	EncodingModeMultiplicator *float64 `json:"encodingModeMultiplicator,omitempty"`
	PerTitleResultStream StatisticsPerTitleStream `json:"perTitleResultStream,omitempty"`
	// The output minutes multiplicator for per-title
	PerTitleMultiplicator *float64 `json:"perTitleMultiplicator,omitempty"`
	PsnrMode PsnrPerStreamMode `json:"psnrMode,omitempty"`
	// The output minutes multiplicator for psnr streams
	PsnrMultiplicator *float64 `json:"psnrMultiplicator,omitempty"`
}

