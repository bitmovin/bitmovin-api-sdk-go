package model

import (
    "encoding/json"
)

// AsperaInput model
type AsperaInput struct {
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
    // Minimal download bandwidth. Examples: 100k, 100m, 100g
    MinBandwidth *string `json:"minBandwidth,omitempty"`
    // Maximal download bandwidth. Examples: 100k, 100m, 100g
    MaxBandwidth *string `json:"maxBandwidth,omitempty"`
    // Host to use for Aspera transfers (required)
    Host *string `json:"host,omitempty"`
    // Username to log into Aspera host (either password and user must be set or token)
    Username *string `json:"username,omitempty"`
    // corresponding password (either password and user must be set or token)
    Password *string `json:"password,omitempty"`
    // Token used for authentication (either password and user must be set or token)
    Token *string `json:"token,omitempty"`
    // Set the TCP port to be used for fasp session initiation
    SshPort *int32 `json:"sshPort,omitempty"`
    // Set the UDP port to be used by fasp for data transfer
    FaspPort *int32 `json:"faspPort,omitempty"`
}

func (m AsperaInput) InputType() InputType {
    return InputType_ASPERA
}
func (m AsperaInput) MarshalJSON() ([]byte, error) {
    type M AsperaInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "ASPERA"

    return json.Marshal(x)
}


