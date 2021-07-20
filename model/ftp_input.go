package model

import (
	"encoding/json"
)

// FtpInput model
type FtpInput struct {
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
	// Ensure that connections originate from the declared ftp host. Default is true.
	RemoteVerificationEnabled *bool `json:"remoteVerificationEnabled,omitempty"`
}

func (m FtpInput) InputType() InputType {
	return InputType_FTP
}
func (m FtpInput) MarshalJSON() ([]byte, error) {
	type M FtpInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "FTP"

	return json.Marshal(x)
}
