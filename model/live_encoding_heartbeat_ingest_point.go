package model

// Details about an individual ingestPoint within the live ingest.
type LiveEncodingHeartbeatIngestPoint struct {
	// The name of the ingestPoint of the original Input resource.
	Name *string `json:"name,omitempty"`
	// Id of the original Input resource. Note that multiple input points (main and backup) can be part of a single Input resource.
	InputId   *string   `json:"inputId,omitempty"`
	InputType InputType `json:"inputType,omitempty"`
	// Indicates whether this particular input is active.
	IsActive *bool `json:"isActive,omitempty"`
	// Indicates whether this particular input is a backup input.
	IsBackup *bool `json:"isBackup,omitempty"`
}
