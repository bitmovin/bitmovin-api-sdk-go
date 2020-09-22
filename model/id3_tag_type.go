package model

// Id3TagType : Id3TagType model
type Id3TagType string

// List of possible Id3TagType values
const (
	Id3TagType_RAW        Id3TagType = "RAW"
	Id3TagType_FRAME_ID   Id3TagType = "FRAME_ID"
	Id3TagType_PLAIN_TEXT Id3TagType = "PLAIN_TEXT"
)
