package model

// PrewarmEncoderSettings model
type PrewarmEncoderSettings struct {
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
	// Encoder Version to be prewarmed. Only one encoder of this version can be prewarmed per cluster. (required)
	EncoderVersion *string `json:"encoderVersion,omitempty"`
	// The minimum number of prewarmed encoders of this Version (required)
	MinPrewarmed *int32 `json:"minPrewarmed,omitempty"`
	// The maximum number of concurrent prewarmed encoders of this Version
	MaxPrewarmed *int32   `json:"maxPrewarmed,omitempty"`
	LogLevel     LogLevel `json:"logLevel,omitempty"`
}
