package model

import (
	"encoding/json"
)

// Live Media and Metadata Ingest Protocol. See https://tools.ietf.org/html/draft-mekuria-mmediaingest-01.
type LiveMediaIngestOutput struct {
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
	Id  *string    `json:"id,omitempty"`
	Acl []AclEntry `json:"acl,omitempty"`
	// URL specifying the publishing point for the output. Can use either http or https. (required)
	PublishingPoint *string `json:"publishingPoint,omitempty"`
}

func (m LiveMediaIngestOutput) OutputType() OutputType {
	return OutputType_LIVE_MEDIA_INGEST
}
func (m LiveMediaIngestOutput) MarshalJSON() ([]byte, error) {
	type M LiveMediaIngestOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "LIVE_MEDIA_INGEST"

	return json.Marshal(x)
}
