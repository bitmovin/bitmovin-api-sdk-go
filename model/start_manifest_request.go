package model

// StartManifestRequest model
type StartManifestRequest struct {
	// Major version of the manifest generator to be used. `V2` is the recommended option and requires the following minimum encoder versions: 2.121.0 for DASH, 2.111.0 for HLS, 2.108.0 for SMOOTH. The default value depends on the sign-up date of your organization. See [documentation](https://developer.bitmovin.com/encoding/docs/manifest-generator-v2) page for a detailed explanation.
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
}
