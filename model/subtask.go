package model
import (
	"time"
)

type Subtask struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Current status
	Status Status `json:"status,omitempty"`
	// Progress in percent
	Progress *int32 `json:"progress,omitempty"`
	// Name of the subtask
	Name string `json:"name,omitempty"`
	// Task specific messages
	Messages []Message `json:"messages,omitempty"`
	// Timestamp when the subtask was created, expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp when the subtask was last updated, expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Timestamp when the subtask was started, expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Timestamp when the subtask status changed to 'QUEUED', expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	QueuedAt *time.Time `json:"queuedAt,omitempty"`
	// Timestamp when the subtask status changed to to 'RUNNING', expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	RunningAt *time.Time `json:"runningAt,omitempty"`
	// Timestamp when the subtask status changed to 'FINISHED', expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Timestamp when the subtask status changed to 'ERROR', expressed in UTC: YYYY-MM-DDThh:mm:ssZ 
	ErrorAt *time.Time `json:"errorAt,omitempty"`
}

