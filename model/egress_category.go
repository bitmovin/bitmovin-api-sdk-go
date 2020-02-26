package model
type EgressCategory string

// List of EgressCategory
const (
	EgressCategory_TRANSFER_RETRY EgressCategory = "TRANSFER_RETRY"
	EgressCategory_TRANSFER EgressCategory = "TRANSFER"
)