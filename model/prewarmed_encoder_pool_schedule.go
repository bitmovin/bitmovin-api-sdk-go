package model

// PrewarmedEncoderPoolSchedule model
type PrewarmedEncoderPoolSchedule struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// The action to be triggered by the schedule (start or stop) (required)
	Action PrewarmedEncoderPoolAction `json:"action,omitempty"`
	// An instant in the future when the specified action should be triggered. Given as UTC expressed in ISO 8601 format (\"YYYY-MM-DDThh:mm:ssZ\"). Seconds will be ignored. (required)
	TriggerDate *DateTime `json:"triggerDate,omitempty"`
}
