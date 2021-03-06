package model

// ConcatenationInputConfiguration model
type ConcatenationInputConfiguration struct {
	// The ID of the input stream to be concatenated. This can be an ingest input stream or a trimming input stream (required)
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Exactly one input stream of a concatenation must have this set to true, which will be used as reference for scaling, aspect ratio, FPS, sample rate, etc.
	IsMain *bool `json:"isMain,omitempty"`
	// A unique integer value that determines concatenation order (required)
	Position *int32 `json:"position,omitempty"`
	// Inserts a padding sequence (black frames and/or silent audio) before the input stream.
	PaddingBefore *PaddingSequence `json:"paddingBefore,omitempty"`
	// Inserts a padding sequence (black frames and/or silent audio) after the input stream.
	PaddingAfter *PaddingSequence `json:"paddingAfter,omitempty"`
	// Specifies the aspect mode that is used when adapting to the main input stream's aspect ratio
	AspectMode AspectMode `json:"aspectMode,omitempty"`
}
