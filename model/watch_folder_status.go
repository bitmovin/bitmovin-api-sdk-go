package model

// WatchFolderStatus : WatchFolderStatus model
type WatchFolderStatus string

// List of possible WatchFolderStatus values
const (
	WatchFolderStatus_STOPPED WatchFolderStatus = "STOPPED"
	WatchFolderStatus_STARTED WatchFolderStatus = "STARTED"
	WatchFolderStatus_ERROR   WatchFolderStatus = "ERROR"
)
