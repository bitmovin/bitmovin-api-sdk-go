package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingManifestsSmoothRepresentationsAPI intermediary API object with no endpoints
type EncodingManifestsSmoothRepresentationsAPI struct {
	apiClient *apiclient.APIClient

	// Mp4 communicates with '/encoding/manifests/smooth/{manifest_id}/representations/mp4' endpoints
	Mp4 *EncodingManifestsSmoothRepresentationsMp4API
}

// NewEncodingManifestsSmoothRepresentationsAPI constructor for EncodingManifestsSmoothRepresentationsAPI that takes options as argument
func NewEncodingManifestsSmoothRepresentationsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothRepresentationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsSmoothRepresentationsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothRepresentationsAPIWithClient constructor for EncodingManifestsSmoothRepresentationsAPI that takes an APIClient as argument
func NewEncodingManifestsSmoothRepresentationsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothRepresentationsAPI {
	a := &EncodingManifestsSmoothRepresentationsAPI{apiClient: apiClient}
	a.Mp4 = NewEncodingManifestsSmoothRepresentationsMp4APIWithClient(apiClient)

	return a
}
