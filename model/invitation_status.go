package model

// InvitationStatus : Status of the invitation
type InvitationStatus string

// List of possible InvitationStatus values
const (
	InvitationStatus_PENDING  InvitationStatus = "PENDING"
	InvitationStatus_ACCEPTED InvitationStatus = "ACCEPTED"
)
