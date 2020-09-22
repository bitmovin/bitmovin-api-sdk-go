package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsTypeAPI communicates with '/encoding/manifests/{manifest_id}/type' endpoints
type EncodingManifestsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsTypeAPI constructor for EncodingManifestsTypeAPI that takes options as argument
func NewEncodingManifestsTypeAPI(options ...apiclient.APIClientOption) (*EncodingManifestsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsTypeAPIWithClient(apiClient), nil
}

// NewEncodingManifestsTypeAPIWithClient constructor for EncodingManifestsTypeAPI that takes an APIClient as argument
func NewEncodingManifestsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsTypeAPI {
	a := &EncodingManifestsTypeAPI{apiClient: apiClient}
	return a
}

// Get Manifest Type
func (api *EncodingManifestsTypeAPI) Get(manifestId string) (*model.ManifestTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.ManifestTypeResponse
	err := api.apiClient.Get("/encoding/manifests/{manifest_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
