package model

import (
    "encoding/json"
)

// ConcatenationInputStream model
type ConcatenationInputStream struct {
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
    // Id of the resource (required)
    Id *string `json:"id,omitempty"`
    // Concatenation configuration for the output of this stream
    Concatenation []ConcatenationInputConfiguration `json:"concatenation,omitempty"`
}

func (m ConcatenationInputStream) InputStreamType() InputStreamType {
    return InputStreamType_CONCATENATION
}
func (m ConcatenationInputStream) MarshalJSON() ([]byte, error) {
    type M ConcatenationInputStream
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "CONCATENATION"

    return json.Marshal(x)
}


