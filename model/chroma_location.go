package model

// ChromaLocation : ChromaLocation model
type ChromaLocation string

// List of possible ChromaLocation values
const (
	ChromaLocation_UNSPECIFIED ChromaLocation = "UNSPECIFIED"
	ChromaLocation_LEFT        ChromaLocation = "LEFT"
	ChromaLocation_CENTER      ChromaLocation = "CENTER"
	ChromaLocation_TOPLEFT     ChromaLocation = "TOPLEFT"
	ChromaLocation_TOP         ChromaLocation = "TOP"
	ChromaLocation_BOTTOMLEFT  ChromaLocation = "BOTTOMLEFT"
	ChromaLocation_BOTTOM      ChromaLocation = "BOTTOM"
)
