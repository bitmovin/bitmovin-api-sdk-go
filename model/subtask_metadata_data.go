package model

// SubtaskMetadataData model
type SubtaskMetadataData struct {
	AvgFramesEncodedPerSecond *float64 `json:"avgFramesEncodedPerSecond,omitempty"`
	BytesEncoded              *int64   `json:"bytesEncoded,omitempty"`
	FramesAnalysed            *int64   `json:"framesAnalysed,omitempty"`
	FramesEncoded             *int64   `json:"framesEncoded,omitempty"`
	RealtimeFactor            *float64 `json:"realtimeFactor,omitempty"`
	Size                      *int64   `json:"size,omitempty"`
}
