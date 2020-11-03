package model

import (
    "encoding/json"
)

// UdpMulticastInput model
type UdpMulticastInput struct {
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
    // Host name or IP address to use (required)
    Host *string `json:"host,omitempty"`
    // Port to use (required)
    Port *int32 `json:"port,omitempty"`
}

func (m UdpMulticastInput) InputType() InputType {
    return InputType_UDP_MULTICAST
}
func (m UdpMulticastInput) MarshalJSON() ([]byte, error) {
    type M UdpMulticastInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "UDP_MULTICAST"

    return json.Marshal(x)
}


