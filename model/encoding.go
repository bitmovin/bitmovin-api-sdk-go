package model

// Encoding model
type Encoding struct {
	// Name of the resource. Can be freely chosen by the user. (required)
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
	// Type of the encoding
	Type EncodingType `json:"type,omitempty"`
	// Timestamp when the encoding was started, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	StartedAt *DateTime `json:"startedAt,omitempty"`
	// Timestamp when the encoding status changed to \"QUEUED\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	QueuedAt *DateTime `json:"queuedAt,omitempty"`
	// Timestamp when the encoding status changed to to \"RUNNING\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	RunningAt *DateTime `json:"runningAt,omitempty"`
	// Timestamp when the encoding status changed to \"FINISHED\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	FinishedAt *DateTime `json:"finishedAt,omitempty"`
	// Timestamp when the encoding status changed to \"ERROR\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ErrorAt *DateTime `json:"errorAt,omitempty"`
	// Progress of the encoding in percent
	Progress    *int32      `json:"progress,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Specify a list of regions which are used in case the preferred region is down. Currently there are several restrictions. - The region has to be specific or AUTO - The region has to be for the same cloud provider as the default one - You can only configure at most 3 fallback regions
	FallbackCloudRegions []CloudRegion `json:"fallbackCloudRegions,omitempty"`
	// Version of the encoder
	EncoderVersion *string `json:"encoderVersion,omitempty"`
	// Define an external infrastructure to run the encoding on. Note If you set this value, the `cloudRegion` must be 'EXTERNAL'.
	InfrastructureId *string                 `json:"infrastructureId,omitempty"`
	Infrastructure   *InfrastructureSettings `json:"infrastructure,omitempty"`
	// Will be set to the encoder version that was actually used for the encoding. This is especially useful when starting an encoding with a version tag like STABLE or BETA.
	SelectedEncoderVersion *string `json:"selectedEncoderVersion,omitempty"`
	// Will be set to the encoding mode that was actually used for the encoding. This is especially useful when starting an encoding with encoding mode STANDARD.
	SelectedEncodingMode EncodingMode `json:"selectedEncodingMode,omitempty"`
	// Contains the region which was selected when cloudregion:AUTO was specified
	SelectedCloudRegion CloudRegion `json:"selectedCloudRegion,omitempty"`
	// The current status of the encoding.
	Status Status `json:"status,omitempty"`
	// You may pass a list of groups associated with this encoding. This will enable you to group results in the statistics resource
	Labels []string `json:"labels,omitempty"`
}
