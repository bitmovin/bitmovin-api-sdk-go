package model


// Policy : The policy for the permissions.
type Policy string

// List of possible Policy values
const (
    Policy_ALLOW Policy = "ALLOW"
    Policy_DENY Policy = "DENY"
)


