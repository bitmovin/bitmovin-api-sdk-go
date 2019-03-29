package model

type ConcatenationInputConfiguration struct {
	// The id of the input stream that should be used for concatenation. Can be either an ingest input stream, or the result of a trimming input stream
	InputStreamId string `json:"inputStreamId,omitempty"`
	// Exactly one input stream of a concatenation must have this set to true, which will be used as reference for scaling, aspect ratio, FPS, sample rate, etc. 
	IsMain *bool `json:"isMain,omitempty"`
	// Position of the stream
	Position *int32 `json:"position,omitempty"`
}

