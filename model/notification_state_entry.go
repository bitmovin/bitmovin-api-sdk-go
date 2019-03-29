package model
import (
	"time"
)

type NotificationStateEntry struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	State NotificationStates `json:"state,omitempty"`
	// Indicate if notification was sent
	Muted *bool `json:"muted,omitempty"`
	// The notification this state belongs to
	NotificationId string `json:"notificationId,omitempty"`
	// Indicate if triggered for specific resource
	ResourceId string `json:"resourceId,omitempty"`
	TriggeredAt *time.Time `json:"triggeredAt,omitempty"`
}

