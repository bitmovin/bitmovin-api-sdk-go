package model

// EgressCategory : EgressCategory model
type EgressCategory string

// List of possible EgressCategory values
const (
	EgressCategory_TRANSFER_RETRY EgressCategory = "TRANSFER_RETRY"
	EgressCategory_TRANSFER       EgressCategory = "TRANSFER"
)
