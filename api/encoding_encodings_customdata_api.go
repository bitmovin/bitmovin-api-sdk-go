package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/customData' endpoints
type EncodingEncodingsCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsCustomdataAPI constructor for EncodingEncodingsCustomdataAPI that takes options as argument
func NewEncodingEncodingsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsCustomdataAPIWithClient constructor for EncodingEncodingsCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsCustomdataAPI {
	a := &EncodingEncodingsCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Encoding Custom Data
func (api *EncodingEncodingsCustomdataAPI) Get(encodingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
