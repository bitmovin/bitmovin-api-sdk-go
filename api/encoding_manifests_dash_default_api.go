package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsDashDefaultAPI communicates with '/encoding/manifests/dash/default' endpoints
type EncodingManifestsDashDefaultAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashDefaultAPI constructor for EncodingManifestsDashDefaultAPI that takes options as argument
func NewEncodingManifestsDashDefaultAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashDefaultAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashDefaultAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashDefaultAPIWithClient constructor for EncodingManifestsDashDefaultAPI that takes an APIClient as argument
func NewEncodingManifestsDashDefaultAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashDefaultAPI {
	a := &EncodingManifestsDashDefaultAPI{apiClient: apiClient}
	return a
}

// Create Default DASH Manifest
// A Default Manifest is the easiest way to create a manifest file. Its contents will be configured automatically, depending on what output your encoding creates (muxings, thumbnails, sprites, subtitles, DRM information). If you need more control, create a Custom Manifest resource instead. See [documentation](https://developer.bitmovin.com/encoding/docs/default-vs-custom-manifest) page for a comparison
func (api *EncodingManifestsDashDefaultAPI) Create(dashManifestDefault model.DashManifestDefault) (*model.DashManifestDefault, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DashManifestDefault
	err := api.apiClient.Post("/encoding/manifests/dash/default", &dashManifestDefault, &responseModel, reqParams)
	return &responseModel, err
}
