package model

import (
    "encoding/json"
)

// FtpOutput model
type FtpOutput struct {
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
    Acl []AclEntry `json:"acl,omitempty"`
    // Host URL or IP of the FTP server (required)
    Host *string `json:"host,omitempty"`
    // Port to use, standard for FTP: 21
    Port *int32 `json:"port,omitempty"`
    // Use passive mode. Default is true.
    Passive *bool `json:"passive,omitempty"`
    // Your FTP Username
    Username *string `json:"username,omitempty"`
    // Your FTP password
    Password *string `json:"password,omitempty"`
    // Controls which transfer version should be used
    TransferVersion TransferVersion `json:"transferVersion,omitempty"`
    // Restrict maximum concurrent connections. Requires at least version 1.1.0.
    MaxConcurrentConnections *int32 `json:"maxConcurrentConnections,omitempty"`
}

func (m FtpOutput) OutputType() OutputType {
    return OutputType_FTP
}
func (m FtpOutput) MarshalJSON() ([]byte, error) {
    type M FtpOutput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "FTP"

    return json.Marshal(x)
}


