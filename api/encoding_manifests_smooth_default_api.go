package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsSmoothDefaultAPI communicates with '/encoding/manifests/smooth/default' endpoints
type EncodingManifestsSmoothDefaultAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsSmoothDefaultAPI constructor for EncodingManifestsSmoothDefaultAPI that takes options as argument
func NewEncodingManifestsSmoothDefaultAPI(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothDefaultAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsSmoothDefaultAPIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothDefaultAPIWithClient constructor for EncodingManifestsSmoothDefaultAPI that takes an APIClient as argument
func NewEncodingManifestsSmoothDefaultAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothDefaultAPI {
	a := &EncodingManifestsSmoothDefaultAPI{apiClient: apiClient}
	return a
}

// Create Default Smooth Streaming Manifest
// A Default Manifest is the easiest way to create a manifest file. Its contents will be configured automatically, depending on what output your encoding creates (muxings, thumbnails, sprites, subtitles, DRM information). If you need more control, create a Custom Manifest resource instead. See [documentation](https://developer.bitmovin.com/encoding/docs/default-vs-custom-manifest) page for a comparison
func (api *EncodingManifestsSmoothDefaultAPI) Create(smoothManifestDefault model.SmoothManifestDefault) (*model.SmoothManifestDefault, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.SmoothManifestDefault
	err := api.apiClient.Post("/encoding/manifests/smooth/default", &smoothManifestDefault, &responseModel, reqParams)
	return &responseModel, err
}
