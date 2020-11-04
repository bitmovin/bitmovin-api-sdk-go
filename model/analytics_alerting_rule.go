package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AnalyticsAlertingRule model
type AnalyticsAlertingRule struct {
	// The id of the alerting rule
	Id *string `json:"id,omitempty"`
	// License key of the alerting rule
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Type of alerting rule
	Type *string `json:"type,omitempty"`
	// Alerting rule name (required)
	Name         *string                        `json:"name,omitempty"`
	Metric       AnalyticsRuleMetric            `json:"metric,omitempty"`
	Filters      []AnalyticsAbstractFilter      `json:"filters,omitempty"`
	Options      *AnalyticsThresholdRuleOptions `json:"options,omitempty"`
	Notification *AnalyticsAlertingNotification `json:"notification,omitempty"`
}

// UnmarshalJSON unmarshals model AnalyticsAlertingRule from a JSON structure
func (m *AnalyticsAlertingRule) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id           *string                        `json:"id"`
		LicenseKey   *string                        `json:"licenseKey"`
		Type         *string                        `json:"type"`
		Name         *string                        `json:"name"`
		Metric       AnalyticsRuleMetric            `json:"metric"`
		Filters      json.RawMessage                `json:"filters"`
		Options      *AnalyticsThresholdRuleOptions `json:"options"`
		Notification *AnalyticsAlertingNotification `json:"notification"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AnalyticsAlertingRule

	result.Id = data.Id
	result.LicenseKey = data.LicenseKey
	result.Type = data.Type
	result.Name = data.Name
	result.Metric = data.Metric
	result.Options = data.Options
	result.Notification = data.Notification

	allOfFilters, err := UnmarshalAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Filters = allOfFilters

	*m = result

	return nil
}
