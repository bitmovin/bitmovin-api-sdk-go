package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsAkamaiNetstorageCustomdataAPI communicates with '/encoding/outputs/akamai-netstorage/{output_id}/customData' endpoints
type EncodingOutputsAkamaiNetstorageCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsAkamaiNetstorageCustomdataAPI constructor for EncodingOutputsAkamaiNetstorageCustomdataAPI that takes options as argument
func NewEncodingOutputsAkamaiNetstorageCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAkamaiNetstorageCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsAkamaiNetstorageCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAkamaiNetstorageCustomdataAPIWithClient constructor for EncodingOutputsAkamaiNetstorageCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsAkamaiNetstorageCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAkamaiNetstorageCustomdataAPI {
	a := &EncodingOutputsAkamaiNetstorageCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Akamai NetStorage Output Custom Data
func (api *EncodingOutputsAkamaiNetstorageCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/akamai-netstorage/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
