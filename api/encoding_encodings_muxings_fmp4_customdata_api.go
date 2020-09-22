package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4CustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4CustomdataAPI constructor for EncodingEncodingsMuxingsFmp4CustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4CustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4CustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4CustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4CustomdataAPI{apiClient: apiClient}
	return a
}

// Get fMP4 muxing Custom Data
func (api *EncodingEncodingsMuxingsFmp4CustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
