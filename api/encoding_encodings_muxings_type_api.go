package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTypeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/{muxing_id}/type' endpoints
type EncodingEncodingsMuxingsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTypeAPI constructor for EncodingEncodingsMuxingsTypeAPI that takes options as argument
func NewEncodingEncodingsMuxingsTypeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTypeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTypeAPIWithClient constructor for EncodingEncodingsMuxingsTypeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTypeAPI {
	a := &EncodingEncodingsMuxingsTypeAPI{apiClient: apiClient}
	return a
}

// Get Muxing type
func (api *EncodingEncodingsMuxingsTypeAPI) Get(encodingId string, muxingId string) (*model.MuxingTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.MuxingTypeResponse
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/{muxing_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
