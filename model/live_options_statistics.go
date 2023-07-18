package model

// LiveOptionsStatistics model
type LiveOptionsStatistics struct {
	Summary *LiveOptionsSummary `json:"summary,omitempty"`
	// Live options statistics aggregated per day (required)
	Breakdown []LiveOptionsBreakdownEntry `json:"breakdown,omitempty"`
}
