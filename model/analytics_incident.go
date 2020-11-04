package model

// AnalyticsIncident model
type AnalyticsIncident struct {
	// Incident id
	Id *string `json:"id,omitempty"`
	// Rule Id
	RuleId *string `json:"ruleId,omitempty"`
	// Start date of the incident
	Start *string `json:"start,omitempty"`
	// End date of the incident
	End *string `json:"end,omitempty"`
	// Recovery state of incident
	IsRecovered *bool `json:"isRecovered,omitempty"`
}
