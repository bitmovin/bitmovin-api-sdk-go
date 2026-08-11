package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersDolbyLoudnessCustomdataAPI communicates with '/encoding/filters/dolby-loudness/{filter_id}/customData' endpoints
type EncodingFiltersDolbyLoudnessCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersDolbyLoudnessCustomdataAPI constructor for EncodingFiltersDolbyLoudnessCustomdataAPI that takes options as argument
func NewEncodingFiltersDolbyLoudnessCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDolbyLoudnessCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersDolbyLoudnessCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDolbyLoudnessCustomdataAPIWithClient constructor for EncodingFiltersDolbyLoudnessCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersDolbyLoudnessCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDolbyLoudnessCustomdataAPI {
	a := &EncodingFiltersDolbyLoudnessCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Dolby Loudness Filter Custom Data
func (api *EncodingFiltersDolbyLoudnessCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/dolby-loudness/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
