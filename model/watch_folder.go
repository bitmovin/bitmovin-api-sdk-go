package model

// WatchFolder model
type WatchFolder struct {
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
	Input      *WatchFolderInput       `json:"input,omitempty"`
	Outputs    []WatchFolderOutput     `json:"outputs,omitempty"`
	// The current status of the Watch Folder. The default value is `STOPPED` (required)
	Status WatchFolderStatus `json:"status,omitempty"`
	// A description text of the current status (required)
	StatusText *string `json:"statusText,omitempty"`
}
