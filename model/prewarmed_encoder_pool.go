package model

// PrewarmedEncoderPool model
type PrewarmedEncoderPool struct {
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
	// The encoder version which the pool's instances will be running (required)
	EncoderVersion *string     `json:"encoderVersion,omitempty"`
	CloudRegion    CloudRegion `json:"cloudRegion,omitempty"`
	// Define an external infrastructure to run the pool on.
	InfrastructureId *string                  `json:"infrastructureId,omitempty"`
	DiskSize         PrewarmedEncoderDiskSize `json:"diskSize,omitempty"`
	// Number of instances to keep prewarmed while the pool is running (required)
	TargetPoolSize *int32                     `json:"targetPoolSize,omitempty"`
	Status         PrewarmedEncoderPoolStatus `json:"status,omitempty"`
}
