package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsPackedAudioCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsPackedAudioCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsPackedAudioCustomdataAPI constructor for EncodingEncodingsMuxingsPackedAudioCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioCustomdataAPI {
	a := &EncodingEncodingsMuxingsPackedAudioCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Packed Audio muxing Custom Data
func (api *EncodingEncodingsMuxingsPackedAudioCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
