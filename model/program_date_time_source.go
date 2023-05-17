package model

// ProgramDateTimeSource : Source for the EXT-X-PROGRAM-DATETIME tag
type ProgramDateTimeSource string

// List of possible ProgramDateTimeSource values
const (
	ProgramDateTimeSource_SYSTEM_CLOCK ProgramDateTimeSource = "SYSTEM_CLOCK"
	ProgramDateTimeSource_EMBEDDED     ProgramDateTimeSource = "EMBEDDED"
)
