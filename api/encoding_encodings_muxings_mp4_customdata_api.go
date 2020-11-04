package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4CustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4CustomdataAPI constructor for EncodingEncodingsMuxingsMp4CustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4CustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4CustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4CustomdataAPI {
	a := &EncodingEncodingsMuxingsMp4CustomdataAPI{apiClient: apiClient}
	return a
}

// Get MP4 muxing Custom Data
func (api *EncodingEncodingsMuxingsMp4CustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
