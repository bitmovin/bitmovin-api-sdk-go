package model

// LiveStandbyPoolResponse model
type LiveStandbyPoolResponse struct {
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
	// Number of instances currently in ready state in the pool
	ReadyEncodings *int32 `json:"readyEncodings,omitempty"`
	// Number of instances currently being prepared in the pool
	PreparingEncodings *int32 `json:"preparingEncodings,omitempty"`
	// Number of instances currently in error state in the pool
	ErrorEncodings *int32 `json:"errorEncodings,omitempty"`
	// The name of the encoding template used with this Standby Pool
	EncodingTemplateName *string               `json:"encodingTemplateName,omitempty"`
	PoolStatus           LiveStandbyPoolStatus `json:"poolStatus,omitempty"`
}
