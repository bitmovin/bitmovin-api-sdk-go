package model

import (
    "encoding/json"
)

// AzureInput model
type AzureInput struct {
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
    // Azure Account Name (required)
    AccountName *string `json:"accountName,omitempty"`
    // Azure Account Key (required)
    AccountKey *string `json:"accountKey,omitempty"`
    // Name of the bucket (required)
    Container *string `json:"container,omitempty"`
}

func (m AzureInput) InputType() InputType {
    return InputType_AZURE
}
func (m AzureInput) MarshalJSON() ([]byte, error) {
    type M AzureInput
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "AZURE"

    return json.Marshal(x)
}


