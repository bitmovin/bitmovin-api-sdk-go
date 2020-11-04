package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsInputStreamsSidecarAPI intermediary API object with no endpoints
type EncodingEncodingsInputStreamsSidecarAPI struct {
	apiClient *apiclient.APIClient

	// DolbyVisionMetadataIngest communicates with '/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest' endpoints
	DolbyVisionMetadataIngest *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPI
}

// NewEncodingEncodingsInputStreamsSidecarAPI constructor for EncodingEncodingsInputStreamsSidecarAPI that takes options as argument
func NewEncodingEncodingsInputStreamsSidecarAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsSidecarAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsSidecarAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsSidecarAPIWithClient constructor for EncodingEncodingsInputStreamsSidecarAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsSidecarAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsSidecarAPI {
	a := &EncodingEncodingsInputStreamsSidecarAPI{apiClient: apiClient}
	a.DolbyVisionMetadataIngest = NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestAPIWithClient(apiClient)

	return a
}
