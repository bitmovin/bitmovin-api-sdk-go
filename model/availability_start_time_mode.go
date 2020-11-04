package model

// AvailabilityStartTimeMode : AvailabilityStartTimeMode model
type AvailabilityStartTimeMode string

// List of possible AvailabilityStartTimeMode values
const (
	AvailabilityStartTimeMode_ON_FIRST_SEGMENT AvailabilityStartTimeMode = "ON_FIRST_SEGMENT"
	AvailabilityStartTimeMode_ON_STREAM_INGEST AvailabilityStartTimeMode = "ON_STREAM_INGEST"
)
