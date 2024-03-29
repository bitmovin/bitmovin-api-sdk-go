package model

// SccCaption model
type SccCaption struct {
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
	// Input location of the SCC file (required)
	Input *InputPath `json:"input,omitempty"`
	// Flavor of SMPTE timecodes in the SCC file (drop-frame or non-drop)
	SmpteTimecodeFlavor SmpteTimecodeFlavor `json:"smpteTimecodeFlavor,omitempty"`
}
