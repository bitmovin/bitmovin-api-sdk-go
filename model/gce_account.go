package model

// GceAccount model
type GceAccount struct {
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
	// Email address of the Google service account that will be used to spin up VMs (required)
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`
	// Google private key of the Google service account that will be used to spin up VMs (required)
	PrivateKey *string `json:"privateKey,omitempty"`
	// ID of the GCP project in which the VMs are supposed to run. (required)
	ProjectId *string `json:"projectId,omitempty"`
}
