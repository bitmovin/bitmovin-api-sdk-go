package model

// EncodingOutputPathsHlsManifest model
type EncodingOutputPathsHlsManifest struct {
	// Id of the HLS manifest
	Id *string `json:"id,omitempty"`
	// Path to the index file of the HLS manifest on the given output
	Path *string `json:"path,omitempty"`
}
