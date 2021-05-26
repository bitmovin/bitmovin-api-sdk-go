package model

import (
	"encoding/json"
)

// AnalyticsAzureOutput model
type AnalyticsAzureOutput struct {
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
	// Azure Account Name (required)
	AccountName *string `json:"accountName,omitempty"`
	// Microsoft Azure Account Access Key with the `Contributor Role` (required)
	AccountKey *string `json:"accountKey,omitempty"`
	// Microsoft Azure container name (required)
	Container *string `json:"container,omitempty"`
}

func (m AnalyticsAzureOutput) AnalyticsOutputType() AnalyticsOutputType {
	return AnalyticsOutputType_AZURE
}
func (m AnalyticsAzureOutput) MarshalJSON() ([]byte, error) {
	type M AnalyticsAzureOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AZURE"

	return json.Marshal(x)
}
