package model
import (
	"time"
)

type EncodingStatisticsLive struct {
	// Date, format. yyyy-MM-dd (required)
	Date *time.Time `json:"date,omitempty"`
	// Bytes encoded for this encoding. (required)
	BytesEncoded *int64 `json:"bytesEncoded,omitempty"`
	// Time in seconds encoded for this encoding. (required)
	TimeEncoded *int64 `json:"timeEncoded,omitempty"`
	// Egress output generated by file transfers in bytes (required)
	BytesEgress *int64 `json:"bytesEgress,omitempty"`
	BillableEncodingMinutes []BillableEncodingMinutes `json:"billableEncodingMinutes,omitempty"`
	BillableEgressBytes []EgressInformation `json:"billableEgressBytes,omitempty"`
	Streams []StatisticsPerStream `json:"streams,omitempty"`
	Muxings []StatisticsPerMuxing `json:"muxings,omitempty"`
	Features []BillableEncodingFeatureMinutes `json:"features,omitempty"`
	// Billable minutes for the muxings.
	BillableTransmuxingMinutes *float64 `json:"billableTransmuxingMinutes,omitempty"`
	// Billable minutes for the features.
	BillableFeatureMinutes *float64 `json:"billableFeatureMinutes,omitempty"`
}

