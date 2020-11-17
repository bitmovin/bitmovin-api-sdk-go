package model

// ExternalIdMode : ExternalIdMode model
type ExternalIdMode string

// List of possible ExternalIdMode values
const (
	ExternalIdMode_CUSTOM    ExternalIdMode = "CUSTOM"
	ExternalIdMode_GLOBAL    ExternalIdMode = "GLOBAL"
	ExternalIdMode_GENERATED ExternalIdMode = "GENERATED"
)
