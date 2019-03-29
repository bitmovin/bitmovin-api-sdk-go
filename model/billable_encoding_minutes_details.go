package model

type BillableEncodingMinutesDetails struct {
	// Only set if resolution information is not present.
	UNKNOWN *int64 `json:"UNKNOWN,omitempty"`
	// Billable minutes for audio. Available if stream is an audio stream.
	AUDIO *int64 `json:"AUDIO,omitempty"`
	// Billable minutes for SD resolutions.
	SD *int64 `json:"SD,omitempty"`
	// Billable minutes for HD resolutions.
	HD *int64 `json:"HD,omitempty"`
	// Billable minutes for UHD resolutions.
	UHD *int64 `json:"UHD,omitempty"`
}

