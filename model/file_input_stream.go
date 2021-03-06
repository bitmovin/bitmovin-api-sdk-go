package model

import (
	"encoding/json"
)

// FileInputStream model
type FileInputStream struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of input (required)
	InputId *string `json:"inputId,omitempty"`
	// Path to file (required)
	InputPath *string             `json:"inputPath,omitempty"`
	FileType  FileInputStreamType `json:"fileType,omitempty"`
}

func (m FileInputStream) InputStreamType() InputStreamType {
	return InputStreamType_FILE
}
func (m FileInputStream) MarshalJSON() ([]byte, error) {
	type M FileInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "FILE"

	return json.Marshal(x)
}
