package model


// SidecarErrorMode : SidecarErrorMode model
type SidecarErrorMode string

// List of possible SidecarErrorMode values
const (
    SidecarErrorMode_FAIL_ON_ERROR SidecarErrorMode = "FAIL_ON_ERROR"
    SidecarErrorMode_CONTINUE_ON_ERROR SidecarErrorMode = "CONTINUE_ON_ERROR"
)


