package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsZixiCustomdataAPI communicates with '/encoding/inputs/zixi/{input_id}/customData' endpoints
type EncodingInputsZixiCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsZixiCustomdataAPI constructor for EncodingInputsZixiCustomdataAPI that takes options as argument
func NewEncodingInputsZixiCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsZixiCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsZixiCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsZixiCustomdataAPIWithClient constructor for EncodingInputsZixiCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsZixiCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsZixiCustomdataAPI {
	a := &EncodingInputsZixiCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Zixi input Custom Data
func (api *EncodingInputsZixiCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/zixi/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
