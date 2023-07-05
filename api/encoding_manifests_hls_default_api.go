package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsHlsDefaultAPI communicates with '/encoding/manifests/hls/default' endpoints
type EncodingManifestsHlsDefaultAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsDefaultAPI constructor for EncodingManifestsHlsDefaultAPI that takes options as argument
func NewEncodingManifestsHlsDefaultAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsDefaultAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsDefaultAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsDefaultAPIWithClient constructor for EncodingManifestsHlsDefaultAPI that takes an APIClient as argument
func NewEncodingManifestsHlsDefaultAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsDefaultAPI {
	a := &EncodingManifestsHlsDefaultAPI{apiClient: apiClient}
	return a
}

// Create Default HLS Manifest
// A Default Manifest is the easiest way to create a manifest file. Its contents will be configured automatically, depending on what output your encoding creates (muxings, thumbnails, sprites, subtitles, DRM information). If you need more control, create a Custom Manifest resource instead. See [documentation](https://developer.bitmovin.com/encoding/docs/default-vs-custom-manifest) page for a comparison
func (api *EncodingManifestsHlsDefaultAPI) Create(hlsManifestDefault model.HlsManifestDefault) (*model.HlsManifestDefault, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.HlsManifestDefault
	err := api.apiClient.Post("/encoding/manifests/hls/default", &hlsManifestDefault, &responseModel, reqParams)
	return &responseModel, err
}
