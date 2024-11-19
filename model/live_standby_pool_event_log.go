package model

// LiveStandbyPoolEventLog model
type LiveStandbyPoolEventLog struct {
	// UUID of the entry
	Id *string `json:"id,omitempty"`
	// Id of the standby_pool associated with the event log
	StandbyPoolId *string `json:"standbyPoolId,omitempty"`
	// (Optional) Id of the standby pool encoding associated with the event
	StandbyPoolEncodingId *string `json:"standbyPoolEncodingId,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *string `json:"createdAt,omitempty"`
	// Short description of the event
	Message *string `json:"message,omitempty"`
	// Detailed description, payloads, hints on how to resolve errors, etc
	Details   *string                     `json:"details,omitempty"`
	EventType LiveStandbyPoolEventLogType `json:"eventType,omitempty"`
}
