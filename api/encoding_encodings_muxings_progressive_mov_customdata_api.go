package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveMovCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveMovCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveMovCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveMovCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveMovCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveMovCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveMovCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Progressive MOV muxing Custom Data
func (api *EncodingEncodingsMuxingsProgressiveMovCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
