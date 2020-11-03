package model


// VideoFormat : VideoFormat model
type VideoFormat string

// List of possible VideoFormat values
const (
    VideoFormat_UNDEFINED VideoFormat = "UNDEFINED"
    VideoFormat_COMPONENT VideoFormat = "COMPONENT"
    VideoFormat_PAL VideoFormat = "PAL"
    VideoFormat_NTSC VideoFormat = "NTSC"
    VideoFormat_SECAM VideoFormat = "SECAM"
    VideoFormat_MAC VideoFormat = "MAC"
)


