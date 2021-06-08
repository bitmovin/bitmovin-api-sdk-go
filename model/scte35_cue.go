package model

// Scte35Cue model
type Scte35Cue struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Cue out duration in seconds. (required)
	CueDuration *float64 `json:"cueDuration,omitempty"`
	// The ids of the manifests to update. If this property is not set, all the manifests tied to the encoding are updated.
	ManifestIds []string `json:"manifestIds,omitempty"`
}
