package model

type StatisticsPerMuxing struct {
	// ID of the stream
	StreamId string `json:"streamId,omitempty"`
	// ID of the muxing
	MuxingId string `json:"muxingId,omitempty"`
	// Multiplier for the encoded minutes. Depends on muxing type.
	Multiplicator *float64 `json:"multiplicator,omitempty"`
	// Encoded bytes.
	EncodedBytes *int64 `json:"encodedBytes,omitempty"`
	// Resulting minutes you will be charged for.
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	MuxingType MuxingType `json:"muxingType,omitempty"`
}

