package model

// SimpleEncodingVodJobOutputArtifact model
type SimpleEncodingVodJobOutputArtifact struct {
	// Name of the artifact. Currently we provide the URL of the HLS manifest with name 'HLS_MANIFEST_URL' and the URL of the DASH manifest with name 'DASH_MANIFEST_URL'
	Name *string `json:"name,omitempty"`
	// A string value described by the 'name' property. Typically this is an absolute URL pointing to a file on the output you specified for the encoding job. The protocol depends on the type of output: \"s3://\" for AWS S3,\"gcs://\" for Google Cloud Storage, \"https://\" for Azure Blob Storage and Akamai NetStorage )
	Value *string `json:"value,omitempty"`
}
