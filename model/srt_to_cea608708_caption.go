package model

// SrtToCea608708Caption model
type SrtToCea608708Caption struct {
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
	// Input location of the SRT file (required)
	Input *InputPath `json:"input,omitempty"`
	// The channel number to embed the CEA subtitles in (required)
	CcChannel Cea608ChannelType `json:"ccChannel,omitempty"`
	// Character encoding of the input SRT file (required)
	CharacterEncoding CaptionCharacterEncoding `json:"characterEncoding,omitempty"`
}
