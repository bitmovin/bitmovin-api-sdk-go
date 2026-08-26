package model

// A single piece of advisory-relevant imagery detected within a shot, for example for regulatory on-screen disclaimers
type ContentAdvisory struct {
	// The kind of advisory-relevant imagery that was detected (required)
	Category AdvisoryCategory `json:"category,omitempty"`
	// The model's own certainty in this detection. Intended to help prioritise human review rather than as a threshold for discarding advisories: detection is tuned to flag uncertain cases rather than miss them, and shots that could not be analysed are reported with LOW confidence (required)
	Confidence AdvisoryConfidence `json:"confidence,omitempty"`
	// A short explanation of what was seen in the shot
	Reason *string `json:"reason,omitempty"`
}
