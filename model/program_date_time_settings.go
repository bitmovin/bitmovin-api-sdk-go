package model

// ProgramDateTimeSettings model
type ProgramDateTimeSettings struct {
	ProgramDateTimeSource    ProgramDateTimeSource     `json:"programDateTimeSource,omitempty"`
	ProgramDateTimePlacement *ProgramDateTimePlacement `json:"programDateTimePlacement,omitempty"`
}
