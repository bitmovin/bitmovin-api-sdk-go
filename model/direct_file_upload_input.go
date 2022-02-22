package model

import (
	"encoding/json"
)

// DirectFileUploadInput model
type DirectFileUploadInput struct {
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
}

func (m DirectFileUploadInput) InputType() InputType {
	return InputType_DIRECT_FILE_UPLOAD
}
func (m DirectFileUploadInput) MarshalJSON() ([]byte, error) {
	type M DirectFileUploadInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DIRECT_FILE_UPLOAD"

	return json.Marshal(x)
}
