package model

// Either height or width is required. It is also possible to set both properties.
type Bif struct {
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
	// Height of one thumbnail
	Height *int32 `json:"height,omitempty"`
	// Width of one thumbnail. Roku recommends a width of 240 for SD and 320 for HD.
	Width *int32 `json:"width,omitempty"`
	// Distance in seconds between a screenshot (required)
	Distance *float64 `json:"distance,omitempty"`
	// Filename of the Bif image. (required)
	Filename *string          `json:"filename,omitempty"`
	Outputs  []EncodingOutput `json:"outputs,omitempty"`
	// Specifies the aspect mode that is used when both height and width are specified Only supported starting with encoder version `2.85.0`.
	AspectMode ThumbnailAspectMode `json:"aspectMode,omitempty"`
}
