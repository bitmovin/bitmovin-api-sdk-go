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
	// Timestamp when the encoding status changed to \"RUNNING\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	RunningAt *DateTime `json:"runningAt,omitempty"`
	// Timestamp when the encoding status changed to 'FINISHED', 'ERROR', 'CANCELED', or 'TRANSFER_ERROR', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ  Note that this timestamp might be inaccurate for encodings which ran prior to the [1.50.0 REST API release](https://bitmovin.com/docs/encoding/changelogs/rest).
	FinishedAt *DateTime `json:"finishedAt,omitempty"`
	// Timestamp when the encoding status changed to 'ERROR', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ  Note that this timestamp is deprecated and is equivalent to finishedAt in case of an 'ERROR'.
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
	// After the encoding has been started, this will contain the encoder version that was actually used. Especially useful when starting an encoding with a version tag like STABLE or BETA.
	SelectedEncoderVersion *string `json:"selectedEncoderVersion,omitempty"`
	// After the encoding has been started, this will contain the encoding mode that was actually used. Especially useful when starting an encoding with encoding mode STANDARD.
	SelectedEncodingMode EncodingMode `json:"selectedEncodingMode,omitempty"`
	// After the encoding has been started, this will contain the cloud region that was actually used. This will differ from cloudRegion if cloudRegion was set to an unspecific region (e.g. 'AUTO')
	SelectedCloudRegion CloudRegion `json:"selectedCloudRegion,omitempty"`
	// The current status of the encoding.
	Status Status `json:"status,omitempty"`
	// You may pass a list of groups associated with this encoding. This will enable you to group results in the statistics resource
	Labels []string `json:"labels,omitempty"`
}
