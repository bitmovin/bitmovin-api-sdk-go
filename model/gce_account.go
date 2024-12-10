package model

// GCE Cloud Connect Account. Configure either by passing a single service account credentials JSON string or by passing the service account email, private key and project ID individually.
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
	// GCP service account credentials JSON
	ServiceAccountCredentials *string `json:"serviceAccountCredentials,omitempty"`
	// Email address of the Google service account that will be used to spin up VMs
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`
	// Google private key of the Google service account that will be used to spin up VMs
	PrivateKey *string `json:"privateKey,omitempty"`
	// ID of the GCP project in which the VMs are supposed to run.
	ProjectId *string `json:"projectId,omitempty"`
}
