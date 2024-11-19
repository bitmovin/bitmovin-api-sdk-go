package model

// LiveStandbyPoolRequest model
type LiveStandbyPoolRequest struct {
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
	// Number of instances to keep ready for streaming while the pool is running (required)
	TargetPoolSize *int32 `json:"targetPoolSize,omitempty"`
	// Base64 encoded template used to start the encodings in the pool (required)
	EncodingTemplate *string `json:"encodingTemplate,omitempty"`
}
