package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsLocalCustomdataAPI communicates with '/encoding/inputs/local/{input_id}/customData' endpoints
type EncodingInputsLocalCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsLocalCustomdataAPI constructor for EncodingInputsLocalCustomdataAPI that takes options as argument
func NewEncodingInputsLocalCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsLocalCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsLocalCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsLocalCustomdataAPIWithClient constructor for EncodingInputsLocalCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsLocalCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsLocalCustomdataAPI {
	a := &EncodingInputsLocalCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Local Input Custom Data
func (api *EncodingInputsLocalCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/local/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
