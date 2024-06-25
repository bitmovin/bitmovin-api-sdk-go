package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWavCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveWavCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveWavCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWavCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWavCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWavCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWavCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveWavCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Progressive Wav muxing Custom Data
func (api *EncodingEncodingsMuxingsProgressiveWavCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
