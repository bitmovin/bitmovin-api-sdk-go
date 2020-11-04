package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// WebhookNotificationWithStreamConditions model
type WebhookNotificationWithStreamConditions struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Notify when condition resolves after it was met
	Resolve *bool `json:"resolve,omitempty"`
	// Specific resource, e.g. encoding id
	ResourceId *string `json:"resourceId,omitempty"`
	// Last time the notification was triggered
	TriggeredAt  *DateTime `json:"triggeredAt,omitempty"`
	Type         *string   `json:"type,omitempty"`
	EventType    *string   `json:"eventType,omitempty"`
	Category     *string   `json:"category,omitempty"`
	ResourceType *string   `json:"resourceType,omitempty"`
	Muted        *bool     `json:"muted,omitempty"`
	// The destination URL where the webhook data is send to (required)
	Url        *string            `json:"url,omitempty"`
	Conditions *AbstractCondition `json:"conditions,omitempty"`
}

// UnmarshalJSON unmarshals model WebhookNotificationWithStreamConditions from a JSON structure
func (m *WebhookNotificationWithStreamConditions) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id           *string         `json:"id"`
		Resolve      *bool           `json:"resolve"`
		ResourceId   *string         `json:"resourceId"`
		TriggeredAt  *DateTime       `json:"triggeredAt"`
		Type         *string         `json:"type"`
		EventType    *string         `json:"eventType"`
		Category     *string         `json:"category"`
		ResourceType *string         `json:"resourceType"`
		Muted        *bool           `json:"muted"`
		Url          *string         `json:"url"`
		Conditions   json.RawMessage `json:"conditions"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result WebhookNotificationWithStreamConditions

	result.Id = data.Id
	result.Resolve = data.Resolve
	result.ResourceId = data.ResourceId
	result.TriggeredAt = data.TriggeredAt
	result.Type = data.Type
	result.EventType = data.EventType
	result.Category = data.Category
	result.ResourceType = data.ResourceType
	result.Muted = data.Muted
	result.Url = data.Url

	allOfConditions, err := UnmarshalAbstractCondition(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Conditions = &allOfConditions

	*m = result

	return nil
}
