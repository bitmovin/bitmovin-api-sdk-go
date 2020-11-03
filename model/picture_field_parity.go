package model


// PictureFieldParity : Specifies which field of an interlaced frame is assumed to be the first one
type PictureFieldParity string

// List of possible PictureFieldParity values
const (
    PictureFieldParity_AUTO PictureFieldParity = "AUTO"
    PictureFieldParity_TOP_FIELD_FIRST PictureFieldParity = "TOP_FIELD_FIRST"
    PictureFieldParity_BOTTOM_FIELD_FIRST PictureFieldParity = "BOTTOM_FIELD_FIRST"
)


