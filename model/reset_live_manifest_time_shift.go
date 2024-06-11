package model

// ResetLiveManifestTimeShift model
type ResetLiveManifestTimeShift struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Determines how many seconds will be left in the manifest after segments are removed. If this is not set, all but one segment will be removed.
	ResidualPeriodInSeconds *float64 `json:"residualPeriodInSeconds,omitempty"`
	// The ids of the manifests to update. If this property is not set, all manifests tied to the encoding are updated.
	ManifestIds []string `json:"manifestIds,omitempty"`
}
