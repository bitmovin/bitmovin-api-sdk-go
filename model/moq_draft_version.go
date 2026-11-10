package model

// MoqDraftVersion : MoQ Transport (MoQT) draft version used on the wire. Set on the output and used for the session with the relay.
type MoqDraftVersion string

// List of possible MoqDraftVersion values
const (
	MoqDraftVersion_DRAFT_16 MoqDraftVersion = "DRAFT_16"
	MoqDraftVersion_DRAFT_18 MoqDraftVersion = "DRAFT_18"
)
