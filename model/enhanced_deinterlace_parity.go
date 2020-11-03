package model


// EnhancedDeinterlaceParity : Specifies which field of an interlaced frame is assumed to be the first one
type EnhancedDeinterlaceParity string

// List of possible EnhancedDeinterlaceParity values
const (
    EnhancedDeinterlaceParity_TOP_FIELD_FIRST EnhancedDeinterlaceParity = "TOP_FIELD_FIRST"
    EnhancedDeinterlaceParity_BOTTOM_FIELD_FIRST EnhancedDeinterlaceParity = "BOTTOM_FIELD_FIRST"
)


