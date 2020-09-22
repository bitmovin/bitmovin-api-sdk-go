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

// Create HLS Manifest Default
func (api *EncodingManifestsHlsDefaultAPI) Create(hlsManifestDefault model.HlsManifestDefault) (*model.HlsManifestDefault, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.HlsManifestDefault
	err := api.apiClient.Post("/encoding/manifests/hls/default", &hlsManifestDefault, &responseModel, reqParams)
	return &responseModel, err
}
