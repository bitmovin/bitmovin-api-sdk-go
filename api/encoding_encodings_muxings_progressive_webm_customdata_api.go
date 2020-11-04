package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Progressive WebM muxing Custom Data
func (api *EncodingEncodingsMuxingsProgressiveWebmCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
