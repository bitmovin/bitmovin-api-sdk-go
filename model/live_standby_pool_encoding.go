package model

// LiveStandbyPoolEncoding model
type LiveStandbyPoolEncoding struct {
	Id         *string `json:"id,omitempty"`
	CreatedAt  *string `json:"createdAt,omitempty"`
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// ID of the encoding that ready for ingest in the standby pool
	EncodingId   *string                              `json:"encodingId,omitempty"`
	Manifests    []LiveStandbyPoolEncodingManifest    `json:"manifests,omitempty"`
	IngestPoints []LiveStandbyPoolEncodingIngestPoint `json:"ingestPoints,omitempty"`
	Status       LiveStandbyPoolEncodingStatus        `json:"status,omitempty"`
}
