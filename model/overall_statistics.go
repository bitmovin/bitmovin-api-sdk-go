package model

type OverallStatistics struct {
	BytesStoredTotal *int64 `json:"bytesStoredTotal,omitempty"`
	BytesTransferredTotal *int64 `json:"bytesTransferredTotal,omitempty"`
	Storages []StorageStatistics `json:"storages,omitempty"`
}

