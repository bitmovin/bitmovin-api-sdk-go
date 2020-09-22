package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsHttpsCustomdataAPI communicates with '/encoding/inputs/https/{input_id}/customData' endpoints
type EncodingInputsHttpsCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsHttpsCustomdataAPI constructor for EncodingInputsHttpsCustomdataAPI that takes options as argument
func NewEncodingInputsHttpsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsHttpsCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsHttpsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsHttpsCustomdataAPIWithClient constructor for EncodingInputsHttpsCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsHttpsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsHttpsCustomdataAPI {
	a := &EncodingInputsHttpsCustomdataAPI{apiClient: apiClient}
	return a
}

// Get HTTPS Custom Data
func (api *EncodingInputsHttpsCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/https/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
