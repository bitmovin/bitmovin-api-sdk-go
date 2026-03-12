package model

// ESAM signal following the SCTE-250 standard
type EsamSignal struct {
	// The offset from the matched signal in ISO 8601 duration format, accurate to milliseconds
	Offset *string `json:"offset,omitempty"`
	// Base64-encoded SCTE-35 binary data to be inserted into the stream (required)
	Binary *string `json:"binary,omitempty"`
}
