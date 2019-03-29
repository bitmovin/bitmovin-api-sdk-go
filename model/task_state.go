package model
type TaskState string

// List of TaskState
const (
	TaskState_ENQUEUED TaskState = "ENQUEUED"
	TaskState_ASSIGNED TaskState = "ASSIGNED"
	TaskState_PREPARED TaskState = "PREPARED"
	TaskState_INPROGRESS TaskState = "INPROGRESS"
	TaskState_FINISHED TaskState = "FINISHED"
	TaskState_ERROR TaskState = "ERROR"
	TaskState_DEQUEUED TaskState = "DEQUEUED"
)