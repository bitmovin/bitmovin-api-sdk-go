package model

// BurnInSubtitleAssa model
type BurnInSubtitleAssa struct {
	// Id of the burn-in ASSA subtitle (required)
	Id *map[string]interface{} `json:"id,omitempty"`
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
	// Character encoding of the ASSA file (required)
	CharacterEncoding CaptionCharacterEncoding `json:"characterEncoding,omitempty"`
	// The input location to get the ASSA file from (required)
	Input *InputPath `json:"input,omitempty"`
}
