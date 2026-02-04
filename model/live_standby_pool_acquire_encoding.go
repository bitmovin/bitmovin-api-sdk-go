package model

// LiveStandbyPoolAcquireEncoding model
type LiveStandbyPoolAcquireEncoding struct {
	// Note: This is not implemented yet and will be ignored if provided. List of ingest points to be used for the acquired encoding. The DNS hostname and RTMPs application name and streamKey will be assigned to the encoding.
	IngestPoints []LiveStandbyPoolEncodingIngestPoint `json:"ingestPoints,omitempty"`
}
