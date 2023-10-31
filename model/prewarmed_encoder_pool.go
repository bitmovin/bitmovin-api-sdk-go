package model

// PrewarmedEncoderPool model
type PrewarmedEncoderPool struct {
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
	// The encoder version which the pool's instances will be running (required)
	EncoderVersion *string `json:"encoderVersion,omitempty"`
	// The cloud region in which the pool's instances will be running. Must be a specific region (e.g. not 'AUTO', 'GOOGLE' or 'EUROPE') (required)
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Define an external infrastructure to run the pool on.
	InfrastructureId *string `json:"infrastructureId,omitempty"`
	// Disk size of the prewarmed instances in GB. Needs to be chosen depending on input file sizes and encoding features used. (required)
	DiskSize PrewarmedEncoderDiskSize `json:"diskSize,omitempty"`
	// Number of instances to keep prewarmed while the pool is running (required)
	TargetPoolSize *int32 `json:"targetPoolSize,omitempty"`
	// Activate dynamic pool behaviour. Pool will increase/decrease based on usage. Minimum pool size is set by targetPoolSize.
	DynamicPool *bool `json:"dynamicPool,omitempty"`
	// Create pool with GPU instances for hardware encoding presets (e.g., VOD_HARDWARE_SHORTFORM).
	GpuEnabled *bool `json:"gpuEnabled,omitempty"`
	// Current status of the pool.
	Status PrewarmedEncoderPoolStatus `json:"status,omitempty"`
}
