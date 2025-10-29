package model

// SubtaskMetadata model
type SubtaskMetadata struct {
	// The timestamp of the metadata record (required)
	Date *DateTime            `json:"date,omitempty"`
	Data *SubtaskMetadataData `json:"data,omitempty"`
}
