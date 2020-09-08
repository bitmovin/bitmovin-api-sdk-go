package model
import (
	"time"
)

type AkamaiNetStorageInput struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Host to use for Akamai NetStorage transfers (required)
	Host string `json:"host,omitempty"`
	// Your Akamai NetStorage Username (required)
	Username string `json:"username,omitempty"`
	// Your Akamai NetStorage password (required)
	Password string `json:"password,omitempty"`
}
func (o AkamaiNetStorageInput) InputType() InputType {
    return InputType_AKAMAI_NETSTORAGE
}

