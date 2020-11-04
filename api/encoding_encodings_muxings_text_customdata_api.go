package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTextCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsTextCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTextCustomdataAPI constructor for EncodingEncodingsMuxingsTextCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsTextCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTextCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTextCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTextCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsTextCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTextCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTextCustomdataAPI {
	a := &EncodingEncodingsMuxingsTextCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Text muxing custom data
func (api *EncodingEncodingsMuxingsTextCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
