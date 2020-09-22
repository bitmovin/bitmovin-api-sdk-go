package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsAkamaiNetstorageCustomdataAPI communicates with '/encoding/inputs/akamai-netstorage/{input_id}/customData' endpoints
type EncodingInputsAkamaiNetstorageCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsAkamaiNetstorageCustomdataAPI constructor for EncodingInputsAkamaiNetstorageCustomdataAPI that takes options as argument
func NewEncodingInputsAkamaiNetstorageCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsAkamaiNetstorageCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsAkamaiNetstorageCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsAkamaiNetstorageCustomdataAPIWithClient constructor for EncodingInputsAkamaiNetstorageCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsAkamaiNetstorageCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAkamaiNetstorageCustomdataAPI {
	a := &EncodingInputsAkamaiNetstorageCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Akamai NetStorage Custom Data
func (api *EncodingInputsAkamaiNetstorageCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/akamai-netstorage/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
