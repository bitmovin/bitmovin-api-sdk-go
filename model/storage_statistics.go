package model

type StorageStatistics struct {
	BytesStored *int64 `json:"bytesStored,omitempty"`
	BytesTransferred *int64 `json:"bytesTransferred,omitempty"`
	Storage string `json:"storage,omitempty"`
}

