package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersAzureSpeechToCaptionsCustomdataAPI communicates with '/encoding/filters/azure-speech-to-captions/{filter_id}/customData' endpoints
type EncodingFiltersAzureSpeechToCaptionsCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPI constructor for EncodingFiltersAzureSpeechToCaptionsCustomdataAPI that takes options as argument
func NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAzureSpeechToCaptionsCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPIWithClient constructor for EncodingFiltersAzureSpeechToCaptionsCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAzureSpeechToCaptionsCustomdataAPI {
	a := &EncodingFiltersAzureSpeechToCaptionsCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Azure Speech to captions Filter Custom Data
func (api *EncodingFiltersAzureSpeechToCaptionsCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/azure-speech-to-captions/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
