package model

import (
	"encoding/json"
)

// SftpOutput model
type SftpOutput struct {
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
	// Deprecation notice: This property does not have any effect and will not be returned by GET endpoints
	Acl []AclEntry `json:"acl,omitempty"`
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
	// Controls which transfer version should be used
	TransferVersion TransferVersion `json:"transferVersion,omitempty"`
	// Restrict maximum concurrent connections. Requires at least version 1.1.0.
	MaxConcurrentConnections *int32 `json:"maxConcurrentConnections,omitempty"`
}

func (m SftpOutput) OutputType() OutputType {
	return OutputType_SFTP
}
func (m SftpOutput) MarshalJSON() ([]byte, error) {
	type M SftpOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "SFTP"

	return json.Marshal(x)
}
