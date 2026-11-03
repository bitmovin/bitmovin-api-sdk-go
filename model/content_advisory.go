package model

// A content advisory reported for a shot, based on detected imagery or a conservative assumption when the analysis was blocked.
type ContentAdvisory struct {
	// Kind of advisory-relevant imagery reported, or conservatively assumed when the shot has status BLOCKED. (required)
	Category AdvisoryCategory `json:"category,omitempty"`
	// Confidence tier derived from the visual evidence reported by the analysis, for prioritizing human review rather than as a calibrated probability. (required)
	Confidence AdvisoryConfidence `json:"confidence,omitempty"`
	// Explanation of the reported visual evidence, or of the conservative assumption when the shot could not be assessed. Retains its original value in language-specific results.
	Reason *string `json:"reason,omitempty"`
}
