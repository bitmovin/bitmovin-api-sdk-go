package model

// SmoothManifestDefault model
type SmoothManifestDefault struct {
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
	Type       ManifestType            `json:"type,omitempty"`
	// The outputs to store the manifest (required)
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Current status
	Status Status `json:"status,omitempty"`
	// Filename of the server manifest
	ServerManifestName *string `json:"serverManifestName,omitempty"`
	// Filename of the client manifest
	ClientManifestName *string `json:"clientManifestName,omitempty"`
	// The id of the encoding to create a default manifest from. (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Specifies the algorithm that determines which output of the given encoding is included into the manifest. Note that this is not related to the \"manifestGenerator\" version of the \"Start\" request.
	Version SmoothManifestDefaultVersion `json:"version,omitempty"`
}
