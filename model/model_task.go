package model

// ModelTask model
type ModelTask struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Current status (required)
	Status Status `json:"status,omitempty"`
	// Estimated ETA in seconds
	Eta *float64 `json:"eta,omitempty"`
	// Progress in percent
	Progress *int32 `json:"progress,omitempty"`
	// List of subtasks
	Subtasks []Subtask `json:"subtasks,omitempty"`
	// Task specific messages
	Messages []Message `json:"messages,omitempty"`
	// Timestamp when the task was created, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Timestamp when the task status changed to \"QUEUED\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	QueuedAt *DateTime `json:"queuedAt,omitempty"`
	// Timestamp when the task status changed to to \"RUNNING\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	RunningAt *DateTime `json:"runningAt,omitempty"`
	// Timestamp when the task status changed to \"FINISHED\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	FinishedAt *DateTime `json:"finishedAt,omitempty"`
	// Timestamp when the task status changed to \"ERROR\", returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ErrorAt *DateTime `json:"errorAt,omitempty"`
	// Additional optional error details
	Error *ErrorDetails `json:"error,omitempty"`
}
