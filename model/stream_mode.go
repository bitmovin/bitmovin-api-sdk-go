package model

// StreamMode : StreamMode model
type StreamMode string

// List of possible StreamMode values
const (
	StreamMode_STANDARD                                        StreamMode = "STANDARD"
	StreamMode_PER_TITLE_TEMPLATE                              StreamMode = "PER_TITLE_TEMPLATE"
	StreamMode_PER_TITLE_TEMPLATE_FIXED_RESOLUTION             StreamMode = "PER_TITLE_TEMPLATE_FIXED_RESOLUTION"
	StreamMode_PER_TITLE_TEMPLATE_FIXED_RESOLUTION_AND_BITRATE StreamMode = "PER_TITLE_TEMPLATE_FIXED_RESOLUTION_AND_BITRATE"
	StreamMode_PER_TITLE_RESULT                                StreamMode = "PER_TITLE_RESULT"
)
