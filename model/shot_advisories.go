package model

// Advisory assessment for a shot.
type ShotAdvisories struct {
	// Outcome of the advisory assessment for this shot. (required)
	Status AdvisoryAnalysisStatus `json:"status,omitempty"`
	// Advisories reported for this shot. Consult status before interpreting an empty or omitted list. (required)
	Advisories []ContentAdvisory `json:"advisories,omitempty"`
}
