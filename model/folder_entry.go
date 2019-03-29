package model

type FolderEntry struct {
	Type FolderEntryType `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

