package model

// StreamInput model
type StreamInput struct {
	// ID of an Input resource defining the input storage. Required if 'inputStreamId' is not set.
	InputId *string `json:"inputId,omitempty"`
	// Path to an input media file. Required if 'inputStreamId' is not set.
	InputPath *string `json:"inputPath,omitempty"`
	// Specifies the strategy for selecting a stream from the input file. Must not be set when 'inputStreamId' is set.
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream to be selected from the input file (zero-based). Must not be set in combination with selectionMode 'AUTO', defaults to 0 for any other selectionMode.
	Position *int32 `json:"position,omitempty"`
	// Set this property instead of all others to reference an InputStream resource (e.g. an Ingest-, Trimming- or ConcatenationInputStream)
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Input analysis details  This property is populated after the encoding has finished
	AnalysisDetails *EncodingStreamInputDetails `json:"analysisDetails,omitempty"`
}
