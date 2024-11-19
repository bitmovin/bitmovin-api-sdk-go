package model

// LiveStandbyPoolEncodingManifest model
type LiveStandbyPoolEncodingManifest struct {
	// URL to the manifest
	Url *string `json:"url,omitempty"`
	// ID of the manifest that was created for the encoding
	ManifestId *string                             `json:"manifestId,omitempty"`
	Type       LiveStandbyPoolEncodingManifestType `json:"type,omitempty"`
}
