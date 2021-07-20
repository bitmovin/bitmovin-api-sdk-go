package model

import (
	"encoding/json"
)

// An external WebVTT file that is added to an encoding. The size limit for a sidecar file is 10 MB
type WebVttSidecarFile struct {
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
	// Path to sidecar file (required)
	InputPath    *string                        `json:"inputPath,omitempty"`
	Outputs      []EncodingOutput               `json:"outputs,omitempty"`
	ErrorMode    SidecarErrorMode               `json:"errorMode,omitempty"`
	Segmentation *WebVttSidecarFileSegmentation `json:"segmentation,omitempty"`
}

func (m WebVttSidecarFile) SidecarFileType() SidecarFileType {
	return SidecarFileType_WEB_VTT
}
func (m WebVttSidecarFile) MarshalJSON() ([]byte, error) {
	type M WebVttSidecarFile
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "WEB_VTT"

	return json.Marshal(x)
}
