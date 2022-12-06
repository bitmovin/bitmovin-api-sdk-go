package model

// StartManifestRequest model
type StartManifestRequest struct {
	// Version of the manifest generation engine to be used. The `V2` option is currently only supported for manifests including resources from a single encoding and is only valid in combination with encoder versions >=  `2.108.0`.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
}
