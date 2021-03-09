package model

// StreamInput model
type StreamInput struct {
	// Id of input
	InputId *string `json:"inputId,omitempty"`
	// Path to media file
	InputPath *string `json:"inputPath,omitempty"`
	// Specifies the algorithm how the stream in the input file will be selected
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream
	Position *int32 `json:"position,omitempty"`
	// Set this property instead of all others to reference an ingest, trimming or concatenation input stream
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Input analysis details  This property is populated after the encoding has finished
	AnalysisDetails *EncodingStreamInputDetails `json:"analysisDetails,omitempty"`
}
