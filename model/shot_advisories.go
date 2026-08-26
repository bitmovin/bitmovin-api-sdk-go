package model

// The result of content advisory detection for a shot, covering both what was detected and whether the shot was assessed at all
type ShotAdvisories struct {
	// Whether and how the shot was assessed for content advisories (required)
	Status AdvisoryAnalysisStatus `json:"status,omitempty"`
	// The advisory-relevant imagery detected in this shot. Empty when the shot was assessed and nothing was found, or when it was not assessed at all (required)
	Advisories []ContentAdvisory `json:"advisories,omitempty"`
}
