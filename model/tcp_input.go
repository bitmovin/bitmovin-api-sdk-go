package model

import (
    "encoding/json"
)

// TcpInput model
type TcpInput struct {
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
}

func (m TcpInput) InputType() InputType {
    return InputType_TCP
}
func (m TcpInput) MarshalJSON() ([]byte, error) {
    type M TcpInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "TCP"

    return json.Marshal(x)
}


