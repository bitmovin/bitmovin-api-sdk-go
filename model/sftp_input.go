package model

import (
    "encoding/json"
)

// SftpInput model
type SftpInput struct {
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
    // Host Url or IP of the SFTP server (required)
    Host *string `json:"host,omitempty"`
    // Port to use, standard for SFTP: 22
    Port *int32 `json:"port,omitempty"`
    // Use passive mode. Default is true.
    Passive *bool `json:"passive,omitempty"`
    // Your SFTP Username
    Username *string `json:"username,omitempty"`
    // Your SFTP password
    Password *string `json:"password,omitempty"`
}

func (m SftpInput) InputType() InputType {
    return InputType_SFTP
}
func (m SftpInput) MarshalJSON() ([]byte, error) {
    type M SftpInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "SFTP"

    return json.Marshal(x)
}


