package model

// StreamsTrimmingStatus : StreamsTrimmingStatus model
type StreamsTrimmingStatus string

// List of possible StreamsTrimmingStatus values
const (
	StreamsTrimmingStatus_UNAVAILABLE StreamsTrimmingStatus = "UNAVAILABLE"
	StreamsTrimmingStatus_AVAILABLE   StreamsTrimmingStatus = "AVAILABLE"
	StreamsTrimmingStatus_STARTED     StreamsTrimmingStatus = "STARTED"
	StreamsTrimmingStatus_ERROR       StreamsTrimmingStatus = "ERROR"
	StreamsTrimmingStatus_FINISHED    StreamsTrimmingStatus = "FINISHED"
)
