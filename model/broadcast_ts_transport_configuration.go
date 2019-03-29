package model

type BroadcastTsTransportConfiguration struct {
	// Output rate in bps. The value zero implies to use minimal rate. The minimal rate leaves approximately 15kbps of null packets in the stream.
	Muxrate *float64 `json:"muxrate,omitempty"`
	// Stop mux on errors. If true, implies halt multiplexing when any error is encountered. If false, errors are ignored and multiplexing continues. Note that the recovery from an error will usually result in an illegal transport stream and artifacts on a decoder.
	StopOnError *bool `json:"stopOnError,omitempty"`
	// If true, prevents adaptation fields with length field equal to zero in video, i.e., zero-length AF. Please note that this condition can only occur when pesAlign for the input stream is set to true.
	PreventEmptyAdaptionFieldsInVideo *bool `json:"preventEmptyAdaptionFieldsInVideo,omitempty"`
	// Program Association Table (PAT) repetition rate per second. Number of PATs per second.
	PatRepetitionRatePerSec *float64 `json:"patRepetitionRatePerSec,omitempty"`
	// Program Map Table (PMT) repetition rate per second. Number of PMTs for each program per second.
	PmtRepetitionRatePerSec *float64 `json:"pmtRepetitionRatePerSec,omitempty"`
}

