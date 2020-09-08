package model
import (
	"time"
)

type TransferRetry struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// The current status of the transfer-retry.
	Status Status `json:"status,omitempty"`
	// Timestamp when the transfer-retry was started, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ 
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Timestamp when the transfer-retry status changed to 'FINISHED', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ 
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Timestamp when the transfer-retry status changed to 'ERROR', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ 
	ErrorAt *time.Time `json:"errorAt,omitempty"`
}

