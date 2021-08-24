package model

// WatchFolderInput model
type WatchFolderInput struct {
	// Input id (required)
	InputId *string `json:"inputId,omitempty"`
	// Path of the directory to monitor (required)
	InputPath *string `json:"inputPath,omitempty"`
}
