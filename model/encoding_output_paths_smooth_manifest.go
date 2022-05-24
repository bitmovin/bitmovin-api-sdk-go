package model

// EncodingOutputPathsSmoothManifest model
type EncodingOutputPathsSmoothManifest struct {
	// Id of the Smooth manifest
	Id *string `json:"id,omitempty"`
	// Path to the client manifest of the Smooth manifest on the given output
	ClientManifestPath *string `json:"clientManifestPath,omitempty"`
	// Path to the server manifest of the Smooth manifest on the given output
	ServerManifestPath *string `json:"serverManifestPath,omitempty"`
}
