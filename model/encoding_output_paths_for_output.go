package model

// EncodingOutputPathsForOutput model
type EncodingOutputPathsForOutput struct {
	// Output paths for Dash manifests
	DashManifests []EncodingOutputPathsDashManifest `json:"dashManifests,omitempty"`
	// Output paths for HLS manifests
	HlsManifests []EncodingOutputPathsHlsManifest `json:"hlsManifests,omitempty"`
	// Output paths for Smooth manifests
	SmoothManifests []EncodingOutputPathsSmoothManifest `json:"smoothManifests,omitempty"`
}
