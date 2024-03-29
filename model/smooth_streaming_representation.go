package model

// SmoothStreamingRepresentation model
type SmoothStreamingRepresentation struct {
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
	// Id of the encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Id of the muxing. (required)
	MuxingId *string `json:"muxingId,omitempty"`
	// The Smooth Streaming ismv or isma file that will be referenced in the manifest. (required)
	MediaFile *string `json:"mediaFile,omitempty"`
	// Language of the MP4 file
	Language *string `json:"language,omitempty"`
	// Track where this MP4 shoudl be added
	TrackName *string `json:"trackName,omitempty"`
	// Specifies the priority of this representation. In the manifest, representations will appear ordered by descending priority values.
	Priority *int32 `json:"priority,omitempty"`
}
