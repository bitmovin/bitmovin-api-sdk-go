package model

// DecodingErrorMode : DecodingErrorMode model
type DecodingErrorMode string

// List of possible DecodingErrorMode values
const (
	DecodingErrorMode_FAIL_ON_ERROR    DecodingErrorMode = "FAIL_ON_ERROR"
	DecodingErrorMode_DUPLICATE_FRAMES DecodingErrorMode = "DUPLICATE_FRAMES"
)
