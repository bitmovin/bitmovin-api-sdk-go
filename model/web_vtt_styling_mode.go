package model

// WebVttStylingMode : Determines how WebVTT styling is handled.
type WebVttStylingMode string

// List of possible WebVttStylingMode values
const (
	WebVttStylingMode_PASSTHROUGH  WebVttStylingMode = "PASSTHROUGH"
	WebVttStylingMode_DROP_STYLING WebVttStylingMode = "DROP_STYLING"
)
