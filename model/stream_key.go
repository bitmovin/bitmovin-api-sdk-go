package model

// StreamKey model
type StreamKey struct {
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
	// Stream key used for live streaming. This stream key is reserved and can be re-used with different live encodings. If this value is not provided, a unique random stream key will be generated. **Important:** This value has to be globally unique. If it is set manually, be sure to use a secure value. If the stream key value is guessed by others your live encoding can be compromised.
	Value *string `json:"value,omitempty"`
	// Status of the stream key
	Status StreamKeyStatus `json:"status,omitempty"`
	// Type of the stream key
	Type StreamKeyType `json:"type,omitempty"`
	// ID of the encoding that is assigned to this stream key. Not set if status is UNASSIGNED
	AssignedEncodingId *string `json:"assignedEncodingId,omitempty"`
	// ID of the ingest point that is assigned to this stream key. Not set if status is UNASSIGNED
	AssignedIngestPointId *string `json:"assignedIngestPointId,omitempty"`
}
