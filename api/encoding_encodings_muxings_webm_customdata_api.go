package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsWebmCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsWebmCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsWebmCustomdataAPI constructor for EncodingEncodingsMuxingsWebmCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsWebmCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmCustomdataAPI {
	a := &EncodingEncodingsMuxingsWebmCustomdataAPI{apiClient: apiClient}
	return a
}

// Get WebM muxing Custom Data
func (api *EncodingEncodingsMuxingsWebmCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
