package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsOutputPathsAPI communicates with '/encoding/encodings/{encoding_id}/output-paths' endpoints
type EncodingEncodingsOutputPathsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsOutputPathsAPI constructor for EncodingEncodingsOutputPathsAPI that takes options as argument
func NewEncodingEncodingsOutputPathsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsOutputPathsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsOutputPathsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsOutputPathsAPIWithClient constructor for EncodingEncodingsOutputPathsAPI that takes an APIClient as argument
func NewEncodingEncodingsOutputPathsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsOutputPathsAPI {
	a := &EncodingEncodingsOutputPathsAPI{apiClient: apiClient}
	return a
}

// Get Encoding Output Paths Retrieval
func (api *EncodingEncodingsOutputPathsAPI) Get(encodingId string) (*[]model.EncodingOutputPaths, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel []model.EncodingOutputPaths
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/output-paths", nil, &responseModel, reqParams)
	return &responseModel, err
}
