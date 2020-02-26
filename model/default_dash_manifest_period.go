package model

type DefaultDashManifestPeriod struct {
	// List the encoding ids for which the conditions should apply
	EncodingIds []string `json:"encodingIds,omitempty"`
	// Adds an adaption set for every item to the period
	AdaptationSets []DefaultManifestCondition `json:"adaptationSets,omitempty"`
}

