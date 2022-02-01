package model

// StartManifestRequest model
type StartManifestRequest struct {
	// Sets the version of the manifest generation engine. The `V2` option is currently only supported for manifests including resources from a single encoding and is only valid in combination with encoder versions >=  `2.108.0`.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
}
