package model

// Subtask model
type Subtask struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Current status (required)
	Status Status `json:"status,omitempty"`
	// Progress in percent
	Progress *int32 `json:"progress,omitempty"`
	// Name of the subtask (required)
	Name *string `json:"name,omitempty"`
	// Task specific messages
	Messages []Message `json:"messages,omitempty"`
	// Task specific metadata
	Metadata []SubtaskMetadata `json:"metadata,omitempty"`
	// Timestamp when the subtask was created, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Timestamp when the subtask was last updated, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	UpdatedAt *DateTime `json:"updatedAt,omitempty"`
	// Timestamp when the subtask was started, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	StartedAt *DateTime `json:"startedAt,omitempty"`
	// Timestamp when the subtask status changed to 'QUEUED', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	QueuedAt *DateTime `json:"queuedAt,omitempty"`
	// Timestamp when the subtask status changed to 'RUNNING', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	RunningAt *DateTime `json:"runningAt,omitempty"`
	// Timestamp when the subtask status changed to a final state like  'FINISHED', 'ERROR', 'CANCELED', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ  Note that this timestamp might be inaccurate for subtasks which ran prior to the [1.50.0 REST API release](https://bitmovin.com/docs/encoding/changelogs/rest).
	FinishedAt *DateTime `json:"finishedAt,omitempty"`
	// Timestamp when the subtask status changed to 'ERROR', returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ  Note that this timestamp is deprecated and is equivalent to finishedAt in case of an 'ERROR'.
	ErrorAt *DateTime `json:"errorAt,omitempty"`
}
